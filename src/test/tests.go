package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	pb "github.com/jmagazine/chat-app-grpc/src/chat"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

func contains(slice []*pb.User, element *pb.User) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}

func TestUserMethods() {
	// Dial connection to grpc server
	if err := godotenv.Load("src/test/.env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	conn, err := grpc.NewClient(os.Getenv("ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := c.DropDatabase(ctx, &pb.DropDatabaseParams{Dbname: "users"}); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "3D000" { // 42P04: database already exists
				fmt.Println("Database does not exist, skipping deletion.")
			} else {
				log.Printf("Failed to drop database: %v", err)
			}
		}
	}

	var new_users = make(map[string]string)
	new_users["Alice"] = "aliceloves123"
	new_users["Bob"] = "bobhates456"
	for fullname, username := range new_users {
		_, err := c.CreateNewUser(ctx, &pb.CreateUserParams{FullName: fullname, Username: username, Password: ""})
		if err != nil {
			log.Fatalf("could not create new user: %v", err)
		}

		params := &pb.GetUsersParams{}
		getUsersRespose, err := c.GetUsers(ctx, params)
		if err != nil {
			log.Fatalf("could not get users: %v", err)
		}

		for _, user := range getUsersRespose.Users {
			if !contains(getUsersRespose.Users, user) {
				out := fmt.Sprint(`
User Details:
Full Name: %s
Username: %s
Id: %d`, user.GetFullName(), user.GetUsername(), user.GetId())
				log.Fatalf(`---Failed TestUserMethods---
User Not Found:
%s`, out)
			}
		}

	}
	fmt.Printf("TestUserMethods - Success\n")

}

func main() {
	TestUserMethods()
}
