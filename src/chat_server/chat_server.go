package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	pb "github.com/jmagazine/chat-app-grpc/src/chat"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	address = "localhost:50051"
)

type ChatServer struct {
	conn *pgx.Conn
	pb.UnimplementedChatServiceServer
}

func NewChatServer() *ChatServer {
	return &ChatServer{}
}

func (server *ChatServer) CreateNewUser(ctx context.Context, in *pb.CreateUserParams) (*pb.User, error) {
	// Protocol to Create new user and update user list
	log.Printf("Received: %v", in.GetFullName())
	createSql := `
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    fullname TEXT,
    username TEXT,
    password TEXT
);`

	_, err := server.conn.Exec(context.Background(), createSql)
	if err != nil {
		fmt.Printf("Failed to create user table: %v\n", err)
		os.Exit(1)
	}
	var new_user = &pb.User{FullName: in.GetFullName(), Username: in.GetUsername(), Password: in.GetPassword()}
	tx, err := server.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
	}
	_, err = tx.Exec(context.Background(), "insert into users(fullname, username, password) values ($1, $2, $3)", new_user.FullName, new_user.Username, new_user.Password)
	if err != nil {
		log.Fatalf("tx.Exec failed: %v", err)
	}

	tx.Commit(context.Background())

	return new_user, nil
}

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

// Attempts to update the user's information, and returns the fields of user after the update. If the update fails,
// it returns the user's fields unedited.
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

func main() {
	// Instantiate database
	if err := godotenv.Load("C:\\Users\\joshm\\GolandProjects\\chat-app-grpc\\src\\.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)

	}
	var password = os.Getenv("DB_PASSWORD")

	dbUrl := fmt.Sprintf("postgres://postgres:%s@localhost:5432/postgres", password)

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	// Instantiate new ChatServer
	var chatServer *ChatServer = NewChatServer()
	chatServer.conn = conn
	if err := chatServer.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
