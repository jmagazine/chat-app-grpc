package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	pb "github.com/jmagazine/chat-app-grpc/src/chat"
	"google.golang.org/grpc"
	"log"
	"net"
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

func (server *ChatServer) Run() error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return grpcServer.Serve(lis)

}

func main() {
	// Instantiate database
	databaseUrl := "postgres://postgres:mysecretpassword@localhost:5432/postgres"
	conn, err := pgx.Connect(context.Background(), databaseUrl)
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
