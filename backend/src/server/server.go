package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	pb "github.com/jmagazine/chat-app-grpc/src/gen/go"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type ChatServer struct {
	sqlDB *pgxpool.Pool
	pb.UnimplementedChatServiceServer
}

func NewChatServer(sqlDB *pgxpool.Pool) ChatServer {
	return ChatServer{sqlDB: sqlDB}
}

// Create the database and users table if they don't exist
func (server ChatServer) RunSQLSetupCommands(ctx context.Context) error {
	createUsersDb := `
CREATE DATABASE IF NOT EXISTS chat-app-grpc;`

	createUsersTbl := `
CREATE TABLE IF NOT EXISTS users (
    Id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    FirstName VARCHAR(255) NOT NULL,
    LastName VARCHAR(255) NOT NULL,
    Username VARCHAR(255) NOT NULL UNIQUE,
    HashToken VARCHAR(65535) NOT NULL
);`

	createUUIDExtension := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
`

	// Create database if it is not present
	_, err := server.sqlDB.Exec(ctx, createUsersDb)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "42P04" { // 42P04: database already exists
				fmt.Println("CreateUser - Database already exists, skipping creation.")
			}
		} else {
			// Other errors
			fmt.Printf("CreateUser - Failed to create user db: %v/n", err)
			return err
		}
	}

	// Create uuid extension if it is not present
	_, err = server.sqlDB.Exec(ctx, createUUIDExtension)
	if err != nil {
		log.Printf("CreateUser - warning: UUIDExtension command was not executed: %v", err)
	}

	// Create table if it is not present
	_, err = server.sqlDB.Exec(ctx, createUsersTbl)
	if err != nil {
		fmt.Printf("CreateUser - Failed to create user table: %v/n", err)
		return err
	}

	return nil
}

// Returns true if the specified value for the specified column in the specified table exists, false otherwise.
func exists(sqlDB *pgxpool.Pool, table string, column string, value string) bool {
	var exists bool
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE %s = $1)`, table, column)
	err := sqlDB.QueryRow(context.Background(), query, value).Scan(&exists)
	if err != nil {
		log.Println("Error checking username existence:", err)
		return false
	}

	return exists
}

// CreateUser defines the protocol to create a new user
func (server ChatServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Printf("Received: %s, %s, %s, %s", in.GetFirstName(), in.GetLastName(), in.GetUsername(), in.GetHashToken())

	err := server.RunSQLSetupCommands(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error.")

	}
	var newUser = &pb.User{FirstName: in.GetFirstName(), LastName: in.GetLastName(), Username: in.GetUsername()}

	tx, err := server.sqlDB.Begin(ctx)
	if err != nil {
		log.Fatalf("CreateUser - conn.Begin failed: %v", err)
		return nil, status.Error(codes.Internal, "Internal server error.")
	}
	defer tx.Rollback(ctx)

	if exists(server.sqlDB, "users", "Username", in.GetUsername()) {
		return nil, status.Error(codes.AlreadyExists, "That username is already taken.")
	}

	// update database
	_, err = tx.Exec(ctx,
		"insert into users(FirstName, LastName, Username, HashToken) values ($1, $2, $3, $4)",
		newUser.FirstName, newUser.LastName, newUser.Username, in.HashToken)
	if err != nil {
		return nil, err
	}

	tx.Commit(ctx)

	return &pb.CreateUserResponse{User: newUser}, nil
}

func (server ChatServer) DeleteAllUsers() error {
	log.Printf("deleting all users")
	ctx := context.Background()
	query := "DELETE FROM users"
	tx, err := server.sqlDB.Begin(ctx)
	if err != nil {
		log.Printf("DeleteAllUsers - failed to begin transaction: %v", err)
		return err
	}

	_, err = tx.Exec(ctx, query)
	if err != nil {
		tx.Rollback(ctx)
		log.Printf("DeleteAllUsers - failed to execute delete command.")
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
		log.Printf("DeleteAllUsers - failed to execute delete command.")
		return err
	}
	log.Printf("deleted all users")
	return nil

}

func (server ChatServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	loginQuery := "SELECT Id, FirstName, LastName, Username FROM users WHERE Username = $1 AND HashToken = $2"
	userRows, err := server.sqlDB.Query(ctx, loginQuery, in.GetUsername(), in.GetHashToken())
	// if querying the database fails, return error code 500
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "User Not Found")
		}
		return nil, status.Error(codes.Internal, "Internal server error.")
	}

	defer userRows.Close()
	user := &pb.User{}
	if !userRows.Next() {
		return nil, status.Error(codes.NotFound, "User Not Found")
	} else {
		if err := userRows.Scan(&user.Id, &user.FirstName, &user.Username); err != nil {
			log.Printf("GetUser - failed to find user in db: %v/n", err)

		}
	}
	return &pb.GetUserResponse{User: user}, nil
}

// DeleteUser deletes a user with the corresponding username from the database if it exists.
func (server ChatServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	user := &pb.User{}
	fmt.Print(in.Username)
	userRows, err := server.sqlDB.Query(ctx, "SELECT Id, FirstName, LastName, Username FROM users WHERE username = $1", in.GetUsername())
	if err != nil {
		return nil, err
	}
	// Check if the query returned a row
	if userRows.Next() {
		// Scan the row into the user struct
		if err := userRows.Scan(&user.Id, &user.FirstName, user.LastName, &user.Username); err != nil {
			log.Printf("DeleteUser - failed to scan user row: %v/n", err)
			return nil, err
		}
	} else {
		log.Printf("DeleteUser - user not found in db: %v/n", err)
		return nil, fmt.Errorf("user not found")
	}

	_, err = server.sqlDB.Exec(ctx, "DELETE FROM users WHERE Username = $1", in.Username)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{User: user}, nil

}

// UpdateUser attempts to update the user's information, and returns the fields of user after the update.
// If the update fails, it returns the user's fields unedited.
func (server ChatServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var builder strings.Builder
	builder.WriteString("UPDATE users SET ")
	allowedKeys := map[string]bool{
		"Id":        true,
		"Username":  true,
		"FirstName": true,
		"LastName":  true,
	}

	updatedFields := in.GetUpdatedFields()
	if len(in.GetUpdatedFields()) > 4 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Too many fields: expected 4, got %d.", len(in.GetUpdatedFields())))
	}

	for key := range updatedFields {
		if !allowedKeys[key] {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Key '%s' not allowed", key))
		}
	}

	values := make([]interface{}, 0, len(in.GetUpdatedFields())+1)
	i := 1
	for k, v := range in.GetUpdatedFields() {
		builder.WriteString(fmt.Sprintf(`%s = $%d, `, k, i))
		values = append(values, v)
		i++
	}

	command := builder.String()
	command = command[:len(command)-2]
	command += fmt.Sprintf(" WHERE Username = $%d", i)
	values = append(values, in.Username)
	tx, err := server.sqlDB.Begin(ctx)
	if err != nil {
		log.Printf("UpdateUser - failed to begin transaction: %v/n", err)
		return nil, err
	}

	// Update the row
	_, err = tx.Exec(ctx, command, values...)
	if err != nil {
		log.Printf("UpdateUser - failed to update user table: %v/n", err)
		tx.Rollback(ctx)
		return nil, err
	}
	tx.Commit(ctx)
	// Get updated user
	row := server.sqlDB.QueryRow(ctx, "SELECT Id, FirstName,  LastName, Username FROM users WHERE Username = $1", in.UpdatedFields["Username"])
	updatedUser := &pb.User{}
	if err := row.Scan(&updatedUser.Id, &updatedUser.FirstName, &updatedUser.Username); err != nil {
		log.Printf("UpdateUser - failed to find updated user: %v/n", err)
		return nil, err
	}
	return &pb.UpdateUserResponse{User: updatedUser}, nil
}

// GetAllUsers returns a list of all registered users.
func (server ChatServer) GetAllUsers(ctx context.Context, in *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	var userSlice []*pb.User
	userRows, err := server.sqlDB.Query(ctx, "select id, FirstName, Username from users")
	if err != nil {
		return nil, err
	}

	defer userRows.Close()
	for userRows.Next() {
		user := pb.User{}
		err = userRows.Scan(&user.Id, &user.FirstName, &user.Username)
		if err != nil {
			return nil, err
		}
		userSlice = append(userSlice, &user)
	}

	return &pb.GetAllUsersResponse{Users: userSlice}, nil
}

// END OF USER FUNCTIONS

// MESSAGE FUNCTIONS

// SendChatMessage sends a chat message to the database.
func (server ChatServer) SendChatMessage(ctx context.Context, in *pb.SendChatMessageRequest) (*pb.SendChatMessageResponse, error) {
	createSql := `CREATE TABLE IF NOT EXISTS chat_messages (
    timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    sender FOREIGN KEY REFERENCES users (username),
    recipient FOREIGN KEY REFERENCES users (username),
    text VARCHAR(10000) NOT NULL DEFAULT "",
);`
	// create db if it doesn't exist
	_, err := server.sqlDB.Exec(ctx, createSql)
	if err != nil {
		log.Printf("conn.Begin failed: %v", err)
	}

	tx, err := server.sqlDB.Begin(ctx)
	if err != nil {
		log.Printf("Failed to begin connection: %v", err)
		return nil, err
	}

	newChatMessage := in.Message
	_, err = tx.Exec(ctx,
		"insert into chat_messages(timestamp, sender, recipient, text) values ($1, $2, $3, $4)",
		newChatMessage.Timestamp, newChatMessage.SenderId, newChatMessage.RecipientId, newChatMessage.Text)
	if err != nil {
		tx.Rollback(ctx)
		log.Printf("SendChatMessage - tx.Exec failed: %v", err)
		return nil, err
	}

	tx.Commit(ctx)
	return &pb.SendChatMessageResponse{Messages: newChatMessage}, nil

}

// GetChatMessages returns all the messages sent between two users.
func (server ChatServer) GetChatMessages(ctx context.Context, in *pb.GetChatMessagesRequest) (*pb.GetChatMessagesResponse, error) {
	getMessagesSql := `SELECT timestamp, sender, recipient FROM chat_messages where sender = $1, recipient = $2`
	_, err := server.sqlDB.Begin(ctx)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	userRows, err := server.sqlDB.Query(ctx, getMessagesSql, in.SenderId, in.RecipientId)
	if err != nil {
		log.Printf("Transaction failed: %v", err)
		return nil, err
	}

	var messageList []*pb.ChatMessage

	for userRows.Next() {
		message := pb.ChatMessage{}
		err = userRows.Scan(&message.Timestamp, &message.SenderId, &message.RecipientId, &message.Text)
		if err != nil {
			return nil, err
		}
		messageList = append(messageList, &message)
	}

	return &pb.GetChatMessagesResponse{Messages: messageList}, nil
}

func (server ChatServer) DropTable(ctx context.Context, in *pb.DropTableRequest) (*pb.DropTableResponse, error) {
	if _, err := server.sqlDB.Exec(ctx, fmt.Sprintf("drop table %s", in.TableName)); err != nil {
		return nil, err
	}
	return &pb.DropTableResponse{}, nil
}

func main() {
	// File path to dev.env
	envPath := filepath.Join("backend", "dev.env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading dev.env file: %v", err)
	}

	var dbURL = os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatalf("DB_URL not specified, quitting with error...")
	}

	// Set up a connection to the chat server.
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		log.Fatalf("HOSTNAME not specified, quitting with error...")
	}

	grpc_port := os.Getenv("GRPC_PORT")
	if grpc_port == "" {
		log.Fatalf("GRPC not specified, quitting with error...")
	}

	chatServiceAddr := hostname + ":" + grpc_port
	conn, err := grpc.NewClient(chatServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to chat service: %v", err)
	}
	fmt.Printf("gRPC server is running on %s:%s\n", hostname, grpc_port)

	sqlDB, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Main - failed to connect to database: %v", err)
	}
	defer conn.Close()
	defer sqlDB.Close()

	// Instantiate new ChatServer
	chatServer := NewChatServer(sqlDB)
	lis, err := net.Listen("tcp", chatServiceAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, chatServer)
	log.Printf("Server listening at %v", lis.Addr())
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Run - failed to serve: %v", err)
		}
	}()

	// Set up CORS options (allow frontend to make calls to api)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins for now, fix later
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Register gRPC server endpoint
	mux := runtime.NewServeMux()
	if err = pb.RegisterChatServiceHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the chat server: %v", err)
	}

	// start listening to requests from the gateway server
	grpc_gateway_port := os.Getenv("GRPC_GATEWAY_PORT")
	if grpc_gateway_port == "" {
		log.Fatalf("GRPC_GATEWAY_PORT not specified, quitting with error...")
	}
	grpc_gateway_addr := fmt.Sprintf("%s:%s", hostname, grpc_gateway_port)
	fmt.Printf("API gateway server is running on %s:%s\n", hostname, grpc_gateway_port)
	if err = http.ListenAndServe(grpc_gateway_addr, corsHandler.Handler(mux)); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}

}
