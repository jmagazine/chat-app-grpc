package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jmagazine/chat-app-grpc/src/utils"

	pb "github.com/jmagazine/chat-app-grpc/src/gen"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.ChatServiceClient
var ctx context.Context
var cancel context.CancelFunc

// Intialize variables before running tests
func InitTests() {

	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	conn, err := grpc.NewClient(os.Getenv("ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client = pb.NewChatServiceClient(conn)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)

	// Close connection
}

func TestUserMethods(ctx context.Context, cancel context.CancelFunc) error {
	// Ensure proper database connection
	AssertClientNotNil()

	hashToken := utils.HashText(os.Getenv("TEST_USER_1_PASSWORD"))

	user1, err := client.CreateUser(ctx, &pb.CreateUserParams{FullName: os.Getenv("TEST_USER_1_FULLNAME"), Username: os.Getenv("TEST_USER_1_USERNAME"), HashToken: hashToken})
	if err != nil {
		return err
	}

	log.Print("Passed: User created successfully.")

	// Ensure attempts to create users with duplicate names fails
	_, err = client.CreateUser(ctx, &pb.CreateUserParams{FullName: os.Getenv("TEST_USER_1_FULLNAME"), Username: os.Getenv("TEST_USER_1_USERNAME"), HashToken: hashToken})

	if err == nil {
		log.Printf("TestUserMethods Failed: Duplicate users were allowed.")
		return err
	}
	log.Printf("Passed: No duplicate users are allowed.")

	// Ensure User was added to db
	user2, err := client.GetUser(ctx, &pb.GetUserParams{Username: user1.GetUsername(), HashToken: utils.HashText(os.Getenv("TEST_USER_1_PASSWORD"))})
	if err != nil {
		log.Printf("TestUserMethods Failed: User was not added to the database: %v", err)
		return err
	}
	// id is handled by the database
	if user1.FullName != user2.FullName ||
		user1.Username != user2.Username {
		log.Printf("TestUserMethods Failed: User fields did not match.")
		return err
	}

	// Ensure fields are properly updated
	var updatedFields = make(map[string]string)
	updatedFields["username"] = "Test User's New Username"
	updatedFields["hash_token"] = utils.HashText("new password")
	newUser, err := client.UpdateUser(ctx, &pb.UpdateUserParams{Username: user2.Username, UpdatedFields: updatedFields})
	if err != nil {
		log.Printf("TestUserMethods Failed: Failed to update user credentials: %v", err)
		return err

	}
	if newUser.Username != updatedFields["username"] {
		log.Printf("TestUserMethods Failed: fields are not consistent.")
		return err

	}
	// Ensure user was deleted
	res, err := client.DeleteUserByUsername(ctx, &pb.DeleteUserByUsernameParams{Username: newUser.Username})
	if err != nil || !res.GetSuccess() {
		log.Printf("TestUserMethods Failed: DeleteUserByUsername failed %v", err)
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

func main() {
	InitTests()

	err := TestUserMethods(ctx, cancel)
	if err != nil {
		log.Printf("TestUserMethods failed: %v", err)
	}

	// Explicitly delete all users, regardless of the test result
	_, delErr := client.DeleteAllUsers(ctx, &pb.DeleteAllUsersParams{})
	if delErr != nil {
		log.Printf("Failed to delete all users: %v", delErr)
	}
}
