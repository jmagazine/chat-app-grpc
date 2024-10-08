package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmagazine/chat-app-grpc/src/utils"

	pb "github.com/jmagazine/chat-app-grpc/src/gen/go"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.ChatServiceClient
var CTX context.Context
var CANCEL context.CancelFunc

func TestUserMethods(CTX context.Context, CANCEL context.CancelFunc) error {
	// Ensure proper database connection
	AssertClientNotNil()

	hashToken := utils.HashText(os.Getenv("TEST_USER_1_PASSWORD"))

	user1Response, err := client.CreateUser(CTX, &pb.CreateUserRequest{FirstName: os.Getenv("TEST_USER_1_FIRST_NAME"), LastName: os.Getenv("TEST_USER_1_LAST_NAME"), Username: os.Getenv("TEST_USER_1_USERNAME"), HashToken: hashToken})
	if err != nil {
		log.Printf("TestUserMethods - failed to create user: %v", err)
		return err
	}

	user1 := user1Response.User
	log.Print("Passed: User created successfully.")

	// Ensure attempts to create users with duplicate names fails
	_, err = client.CreateUser(CTX, &pb.CreateUserRequest{FirstName: os.Getenv("TEST_USER_1_FIRST_NAME"), LastName: os.Getenv("TEST_USER_1_LAST_NAME"), Username: os.Getenv("TEST_USER_1_USERNAME"), HashToken: hashToken})

	if err == nil {
		return errors.New("duplicate users were allowed")
	}

	log.Printf("Passed: No duplicate users are allowed")

	// Ensure User was added to db
	user2Response, err := client.GetUser(CTX, &pb.GetUserRequest{Username: user1.GetUsername(), HashToken: hashToken})
	if err != nil {
		log.Printf("User was not added to the database: %v", err)
		return err
	}

	user2 := user2Response.User
	// id is handled by the database
	if user1.FirstName != user2.FirstName ||
		user1.LastName != user2.LastName ||
		user1.Username != user2.Username {
		log.Printf("User fields did not match.")
		return err
	}

	// Ensure fields are properly updated
	var updatedFields = make(map[string]string)
	updatedFields["Username"] = "Test User 2 New Username"
	updatedFields["HashToken"] = utils.HashText("new password")
	newUserRequest, err := client.UpdateUser(CTX, &pb.UpdateUserRequest{Username: user1.Username, UpdatedFields: updatedFields})
	if err != nil {
		log.Printf("TestUserMethods - Failed to update user credentials: %v", err)
		return err

	}

	newUser := newUserRequest.User
	if newUser.Username != updatedFields["Username"] {
		log.Printf("TestUserMethods - fields are not consistent.")
		return err

	}

	// Ensure user was deleted
	_, err = client.DeleteUser(CTX, &pb.DeleteUserRequest{Username: newUser.Username})
	if err != nil {
		log.Printf("TestUserMethods - DeleteUser failed %v", err)
		return err

	}

	log.Printf("TestUserMethods: Passed!")
	return nil
}

func AssertClientNotNil() {
	if client == nil {
		log.Fatalf("ChatServer was never started. Call InitTests() first.")
	}
}

// Intialize variables before running tests
func InitTests() {

	if err := godotenv.Load("dev.env"); err != nil {
		log.Fatalf("Error loading dev.env file: %v", err)
	}

	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		log.Fatalf("HOSTNAME not specified, quitting with error...")
	}
	grpc_port := os.Getenv("GRPC_PORT")
	if hostname == "" {
		log.Fatalf("GRPC_PORT not specified, quitting with error...")
	}

	addr := fmt.Sprintf("%s:%s", hostname, grpc_port)

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client = pb.NewChatServiceClient(conn)
	CTX, CANCEL = context.WithTimeout(context.Background(), time.Second)

	// Close connection
}
func main() {
	InitTests()

	err := TestUserMethods(CTX, CANCEL)
	if err != nil {
		log.Printf("TestUserMethods - %v", err)
	}

	// Explicitly delete all users, regardless of the test result
	_, delErr := client.DeleteAllUsers(CTX, &pb.DeleteAllUsersRequest{})
	if delErr != nil {
		log.Printf("Failed to delete all users: %v", delErr)
	}
}
