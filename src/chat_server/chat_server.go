package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	pb "github.com/jmagazine/chat-app-grpc/src/chat"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
)

var address string

type ChatServer struct {
	test bool
	conn *pgx.Conn
	pb.UnimplementedChatServiceServer
}

func NewChatServer(test bool) *ChatServer {
	address = os.Getenv("ADDRESS")
	return &ChatServer{test: test}
}

// USER FUNCTIONS

// CreateNewUser defines the protocol to create a new user
func (server *ChatServer) CreateNewUser(ctx context.Context, in *pb.CreateUserParams) (*pb.User, error) {
	log.Printf("Received: %v", in.GetFullName())
	createUsersDb := `
CREATE DATABASE users_%sdb);`

	createUsersTbl := `
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    fullname VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL
);`

	createUUIDExtension := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
`

	// Create database if it is not present
	_, err := server.conn.Exec(context.Background(), createUsersDb)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "42P04" { // 42P04: database already exists
				fmt.Println("Database already exists, skipping creation.")
			}
		} else {
			// Other errors
			fmt.Printf("Failed to create user db: %v\n", err)
			return nil, err
		}
	}
	// Create table if it is not present
	_, err = server.conn.Exec(context.Background(), createUsersTbl)
	if err != nil {
		fmt.Printf("Failed to create user table: %v\n", err)
		return nil, err
	}

	// Create uuid extension if it is not present
	_, err = server.conn.Exec(context.Background(), createUUIDExtension)
	if err != nil {
		fmt.Printf("Failed to create uuid-usop extension: %v\n", err)
		return nil, err
	}
	var new_user = &pb.User{FullName: in.GetFullName(), Username: in.GetUsername(), Password: in.GetPassword()}
	tx, err := server.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
		return nil, err
	}
	// update database
	_, err = tx.Exec(context.Background(),
		"insert into users(fullname, username, password) values ($1, $2, $3)",
		new_user.FullName, new_user.Username, new_user.Password)
	if err != nil {
		tx.Rollback(context.Background())
		log.Printf("tx.Exec failed: %v", err)
		return nil, err
	}

	tx.Commit(context.Background())

	return new_user, nil
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
		log.Fatalf("failed to serve: %v", err)
	}
	return grpcServer.Serve(lis)

}

// DeleteUser deletes a user from the database if it exists.
func (server *ChatServer) DeleteUser(ctx context.Context, in *pb.DeleteUserByIdParams) (*pb.DidDeleteUserMessage, error) {
	// Get user to delete
	rows, err := server.conn.Query(context.Background(), "delete from users where id = $1", in.Id)
	if err != nil {
		message := fmt.Sprintf("%v", err)
		return &pb.DidDeleteUserMessage{Id: in.Id, Success: false, Error: &message}, nil
	}
	defer rows.Close()
	return &pb.DidDeleteUserMessage{Id: in.Id, Success: true}, nil

}

// UpdateUser attempts to update the user's information, and returns the fields of user after the update.
// If the update fails, it returns the user's fields unedited.
func (server *ChatServer) UpdateUser(ctx context.Context, in *pb.UpdateUserParams) (*pb.User, error) {
	var builder strings.Builder
	builder.WriteString("UPDATE users SET ")
	values := make([]interface{}, 0, len(in.GetUpdatedFields()))
	i := 1
	for k, v := range in.GetUpdatedFields() {
		builder.WriteString(fmt.Sprintf(`"%s" = $%d, `, k, i))
		values = append(values, v)
		i++
	}
	builder.WriteString(" WHERE id = $1")
	command := builder.String()
	tx, err := server.conn.Begin(context.Background())

	// Update the row
	_, err = tx.Exec(context.Background(), command, in.Id)
	if err != nil {
		log.Printf("Failed to update user table: %v\n", err)
		return nil, err
	}
	tx.Commit(context.Background())
	// Get updated user
	row := server.conn.QueryRow(context.Background(), "select * from users where id = $1", in.Id)
	updatedUser := &pb.User{}
	if err := row.Scan(&updatedUser.Id, &updatedUser.FullName, &updatedUser.Username, &updatedUser.Password); err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// GetUsers returns a list of all registered users.
func (server *ChatServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UsersList, error) {

	var users_list *pb.UsersList = &pb.UsersList{}
	rows, err := server.conn.Query(context.Background(), "select * from users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		user := pb.User{}
		err = rows.Scan(&user.Id, &user.FullName, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users_list.Users = append(users_list.Users, &user)
	}

	return users_list, nil
}

// END OF USER FUNCTIONS

// MESSAGE FUNCTIONS

// SendChatMessage sends a chat message to the database.
func (server *ChatServer) SendChatMessage(ctx context.Context, in *pb.SendChatMessageParams) (*pb.ChatMessage, error) {
	createSql := `CREATE TABLE IF NOT EXISTS chat_messages (
    timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    sender FOREIGN KEY REFERENCES users (id),
    recipient FOREIGN KEY REFERENCES users (id),
);`
	// create db if it doesn't exist
	_, err := server.conn.Exec(context.Background(), createSql)
	if err != nil {
		log.Printf("conn.Begin failed: %v", err)
	}

	tx, err := server.conn.Begin(context.Background())

	newChatMessage := in.Message
	_, err = tx.Exec(context.Background(),
		"insert into chat_messages(timestamp, sender_id, recipient_id) values ($1, $2, $3)",
		newChatMessage.Timestamp, newChatMessage.SenderId, newChatMessage.RecipientId)
	if err != nil {
		tx.Rollback(context.Background())
		log.Printf("tx.Exec failed: %v", err)
		return nil, err
	}

	tx.Commit(context.Background())
	return newChatMessage, nil

}

func (server *ChatServer) DropDatabase(ctx context.Context, in *pb.DropDatabaseParams) (*pb.DropDatabaseMessage, error) {
	if _, err := server.conn.Exec(context.Background(), fmt.Sprintf("drop database %s", in.Dbname)); err != nil {
		log.Printf("conn.Exec failed: %v", err)
		return &pb.DropDatabaseMessage{Success: false}, err
	}
	return &pb.DropDatabaseMessage{Success: true}, nil
}

func main() {
	// Instantiate database
	if err := godotenv.Load("src/.env.prod"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)

	}
	var dbUrl = os.Getenv("DB_URL")

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	// Instantiate new ChatServer
	var chatServer = NewChatServer(false)
	chatServer.conn = conn
	if err := chatServer.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
