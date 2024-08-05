package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	pb "github.com/jmagazine/chat-app-grpc/src/gen"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var address string

type ChatServer struct {
	conn *pgx.Conn
	pb.UnimplementedChatServiceServer
}

func NewChatServer(test bool) *ChatServer {
	address = os.Getenv("ADDRESS")
	return &ChatServer{}
}

// USER FUNCTIONS

// CreateUser defines the protocol to create a new user
func (server *ChatServer) CreateUser(ctx context.Context, in *pb.CreateUserParams) (*pb.User, error) {
	log.Printf("CreateUser - Received: %v", in.GetFullName())
	createUsersDb := `
CREATE DATABASE chat-app-grpc);`

	createUsersTbl := `
CREATE TABLE IF NOT EXISTS users (
    Id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    fullname VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    hash_token VARCHAR(65535) NOT NULL
);`

	createUUIDExtension := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
`

	// Create database if it is not present
	_, err := server.conn.Exec(ctx, createUsersDb)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "42P04" { // 42P04: database already exists
				fmt.Println("CreateUser - Database already exists, skipping creation.")
			}
		} else {
			// Other errors
			fmt.Printf("CreateUser - Failed to create user db: %v\n", err)
			return nil, err
		}
	}

	// Create uuid extension if it is not present
	_, err = server.conn.Exec(ctx, createUUIDExtension)
	if err != nil {
		log.Printf("CreateUser - warning: UUIDExtension command was not executed: %v", err)
	}

	// Create table if it is not present
	_, err = server.conn.Exec(ctx, createUsersTbl)
	if err != nil {
		fmt.Printf("CreateUser - Failed to create user table: %v\n", err)
		return nil, err
	}

	var new_user = &pb.User{FullName: in.GetFullName(), Username: in.GetUsername()}

	tx, err := server.conn.Begin(ctx)
	if err != nil {
		log.Fatalf("CreateUser - conn.Begin failed: %v", err)
		return nil, err
	}
	// update database
	_, err = tx.Exec(ctx,
		"insert into users(fullname, username, hash_token) values ($1, $2, $3)",
		new_user.FullName, new_user.Username, in.HashToken)
	if err != nil {
		tx.Rollback(ctx)
		log.Printf("CreateUser - tx.Exec failed: %v", err)
		return nil, err
	}

	tx.Commit(ctx)

	return new_user, nil
}

func (server *ChatServer) DeleteAllUsers(ctx context.Context, in *pb.DeleteAllUsersParams) (*pb.DidDeleteAllUsers, error) {
	log.Printf("deleting all users")
	query := "DELETE FROM users"
	tx, err := server.conn.Begin(ctx)

	if err != nil {
		tx.Rollback(ctx)
		log.Printf("DeleteAllUsers - failed to begin transaction.")
		return &pb.DidDeleteAllUsers{}, err
	}

	_, err = tx.Exec(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		log.Printf("DeleteAllUsers - failed to execute delete command.")
		return &pb.DidDeleteAllUsers{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		log.Printf("DeleteAllUsers - failed to execute delete command.")
		return &pb.DidDeleteAllUsers{}, err
	}
	log.Printf("deleted all users")
	return &pb.DidDeleteAllUsers{}, nil

}

func (server *ChatServer) GetUser(ctx context.Context, in *pb.GetUserParams) (*pb.User, error) {

	loginQuery := "SELECT id, fullname, username FROM users WHERE username = $1 AND hash_token = $2"
	rows, err := server.conn.Query(ctx, loginQuery, in.GetUsername(), in.GetHashToken())
	if err != nil {
		log.Printf("GetUser - failed to find user with specified credentials, %v", err)
	}

	defer rows.Close()
	user := &pb.User{}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.FullName, &user.Username); err != nil {
			log.Printf("GetUser - failed to find user in db: %v\n", err)
			return nil, err
		}
	}
	return user, nil
}

// Run starts the server.
func (server *ChatServer) Run() error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, server)
	log.Printf("Server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Run - failed to serve: %v", err)
	}
	return grpcServer.Serve(lis)

}

// DeleteUser deletes a user with the corresponding username from the database if it exists.
func (server *ChatServer) DeleteUserByUsername(ctx context.Context, in *pb.DeleteUserByUsernameParams) (*pb.DidDeleteUserMessage, error) {
	// Get user to delete
	rows, err := server.conn.Query(ctx, "delete from users where username = $1", in.Username)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		return &pb.DidDeleteUserMessage{Success: false, Error: &message}, nil
	}
	defer rows.Close()
	return &pb.DidDeleteUserMessage{Success: true}, nil

}

// UpdateUser attempts to update the user's information, and returns the fields of user after the update.
// If the update fails, it returns the user's fields unedited.
func (server *ChatServer) UpdateUser(ctx context.Context, in *pb.UpdateUserParams) (*pb.User, error) {
	var builder strings.Builder
	builder.WriteString("UPDATE users SET ")
	values := make([]interface{}, 0, len(in.GetUpdatedFields())+1)
	i := 1
	for k, v := range in.GetUpdatedFields() {
		builder.WriteString(fmt.Sprintf(`%s = $%d, `, k, i))
		values = append(values, v)
		i++
	}

	command := builder.String()
	command = command[:len(command)-2]
	command += fmt.Sprintf(" WHERE username = $%d", i)
	values = append(values, in.Username)
	tx, err := server.conn.Begin(ctx)
	if err != nil {
		log.Printf("UpdateUser - failed to begin transaction: %v\n", err)
		return nil, err
	}

	// Update the row
	_, err = tx.Exec(ctx, command, values...)
	if err != nil {
		log.Printf("UpdateUser - failed to update user table: %v\n", err)
		tx.Rollback(ctx)
		return nil, err
	}
	tx.Commit(ctx)
	// Get updated user
	row := server.conn.QueryRow(ctx, "SELECT id, fullname, username FROM users WHERE username = $1", in.UpdatedFields["username"])
	updatedUser := &pb.User{}
	if err := row.Scan(&updatedUser.Id, &updatedUser.FullName, &updatedUser.Username); err != nil {
		log.Printf("UpdateUser - failed to find updated user: %v\n", err)

		return nil, err
	}
	return updatedUser, nil
}

// GetAllUsers returns a list of all registered users.
func (server *ChatServer) GetAllUsers(ctx context.Context, in *pb.GetAllUsersParams) (*pb.UsersList, error) {
	var usersList = &pb.UsersList{}
	rows, err := server.conn.Query(ctx, "select id, fullname, username from users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		user := pb.User{}
		err = rows.Scan(&user.Id, &user.FullName, &user.Username)
		if err != nil {
			return nil, err
		}
		usersList.Users = append(usersList.Users, &user)
	}

	return usersList, nil
}

// END OF USER FUNCTIONS

// MESSAGE FUNCTIONS

// SendChatMessage sends a chat message to the database.
func (server *ChatServer) SendChatMessage(ctx context.Context, in *pb.SendChatMessageParams) (*pb.ChatMessage, error) {
	createSql := `CREATE TABLE IF NOT EXISTS chat_messages (
    timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    sender FOREIGN KEY REFERENCES users (username),
    recipient FOREIGN KEY REFERENCES users (username),
    text VARCHAR(10000) NOT NULL DEFAULT "",
);`
	// create db if it doesn't exist
	_, err := server.conn.Exec(ctx, createSql)
	if err != nil {
		log.Printf("conn.Begin failed: %v", err)
	}

	tx, err := server.conn.Begin(ctx)
	if err != nil {
		log.Printf("Failed to begin connection: %v", err)
		return nil, err
	}

	newChatMessage := in.Message
	_, err = tx.Exec(ctx,
		"insert into chat_messages(timestamp, sender, recipient, text) values ($1, $2, $3, $4)",
		newChatMessage.Timestamp, newChatMessage.Sender, newChatMessage.Recipient, newChatMessage.Text)
	if err != nil {
		tx.Rollback(ctx)
		log.Printf("SendChatMessage - tx.Exec failed: %v", err)
		return nil, err
	}

	tx.Commit(ctx)
	return newChatMessage, nil

}

// GetChatMessages returns all the messages sent between two users.
func (server *ChatServer) GetChatMessages(ctx context.Context, in *pb.GetChatMessagesParams) (*pb.ChatMessageList, error) {
	getMessagesSql := `SELECT timestamp, sender, recipient FROM chat_messages where sender = $1, recipient = $2`
	_, err := server.conn.Begin(ctx)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	rows, err := server.conn.Query(ctx, getMessagesSql, in.SenderId, in.RecipientId)
	if err != nil {
		log.Printf("Transaction failed: %v", err)
		return nil, err
	}

	var messageList = &pb.ChatMessageList{}

	for rows.Next() {
		message := pb.ChatMessage{}
		err = rows.Scan(&message.Timestamp, &message.Sender, &message.Recipient, &message.Text)
		if err != nil {
			return nil, err
		}
		messageList.Messages = append(messageList.Messages, &message)
	}

	return messageList, nil
}

func (server *ChatServer) DropTable(ctx context.Context, in *pb.DropTableParams) (*pb.DropTableMessage, error) {
	if _, err := server.conn.Exec(ctx, fmt.Sprintf("drop table %s", in.TableName)); err != nil {
		log.Printf("conn.Exec failed: %v", err)
		return &pb.DropTableMessage{Success: false}, err
	}
	return &pb.DropTableMessage{Success: true}, nil
}

func main() {
	// Instantiate database
	ctx := context.Background()
	if err := godotenv.Load("../env.dev"); err != nil {
		log.Fatalf("Error loading .env.dev file: %v", err)

	}
	var dbUrl = os.Getenv("DB_URL")

	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Main - failed to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	// Instantiate new ChatServer
	var chatServer = NewChatServer(false)
	chatServer.conn = conn
	if err := chatServer.Run(); err != nil {
		log.Fatalf("Main - failed to serve: %v", err)
	}
}
