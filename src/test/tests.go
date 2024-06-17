package main

import (
	"context"
	pb "github.com/jmagazine/chat-app-grpc/src/chat"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

var client pb.ChatServiceClient
var conn *grpc.ClientConn

func contains(slice []*pb.User, element *pb.User) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}
func InitTests() {
	if err := godotenv.Load("src/test/.env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	conn, err := grpc.NewClient(os.Getenv("ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = pb.NewChatServiceClient(conn)
}

func createUserWithParams(fullName string, userName string, password string) (*pb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.CreateNewUser(ctx, &pb.CreateUserParams{FullName: fullName, Username: userName, Password: password})
	if err != nil {
		return nil, err
	}
	return r, nil
}

// NoDuplicateUsersTest ensures you cannot create users with duplicate usernames.

func TestUserMethods() {
	AssertClientNotNil()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Test if a user is created
	user1 := &pb.User{FullName: time.DateTime, Username: "Test User", Password: "Test User"}
	_, err := createUserWithParams(user1.FullName, user1.Username, user1.Password)
	if err != nil {
		log.Fatalf("TestUserMethods Failed: CreateUserWithParams failed: %v", err)
	}

	// Ensure attempts to create users with duplicate names fails
	_, err = createUserWithParams(user1.FullName, user1.Username, user1.Password)
	if err == nil {
		log.Fatalf("TestUserMethods Failed: Duplicate users were allowed.")
	}

	// Ensure User was added to db
	user2, err := client.GetUserByUsername(ctx, &pb.GetUserByUsernameParams{Username: user1.GetUsername()})
	if err != nil {
		log.Fatalf("TestUserMethods Failed: User was not added to the database: %v", err)
	}
	// id is handled by the database
	if user1.FullName != user2.FullName ||
		user1.Username != user2.Username ||
		user1.Password != user2.Password {
		log.Fatalf("TestUserMethods Failed: User fields did not match.")
	}

	// Ensure fields are properly updated
	var updatedFields = make(map[string]string)
	updatedFields["username"] = "Test User's New Username"
	updatedFields["password"] = "12345678"
	newUser, err := client.UpdateUser(ctx, &pb.UpdateUserParams{Username: user2.Username, UpdatedFields: updatedFields})
	if err != nil {
		log.Fatalf("TestUserMethods Failed: Failed to update user credentials: %v", err)
	}
	if newUser.Username != updatedFields["username"] || newUser.Password != updatedFields["password"] {
		log.Fatalf("TestUserMethods Failed: fields are not consistent.")
	}
	// Ensure user was deleted
	res, err := client.DeleteUserByUsername(ctx, &pb.DeleteUserByUsernameParams{Username: newUser.Username})
	if err != nil || !res.GetSuccess() {
		log.Fatalf("TestUserMethods Failed: DeleteUserByUsername failed %v", err)
	}

	log.Printf("TestUserMethods: Passed!")

}
func AssertClientNotNil() {
	if client == nil {
		log.Fatalf("ChatServer was never started. Call InitTests() first.")
	}
}

func main() {
	InitTests()
	TestUserMethods()
}
