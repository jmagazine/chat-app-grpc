package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jmagazine/chat-app-grpc/utils"

	pb "github.com/jmagazine/chat-app-grpc/chat_server/gen/github.com/chat-app-grpc"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.ChatServiceClient

func InitTests() {
	if err := godotenv.Load("tests/.env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	conn, err := grpc.NewClient(os.Getenv("ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = pb.NewChatServiceClient(conn)
}

func createUserWithParams(ctx *context.Context, fullName string, userName string, hashToken string) (*pb.User, error) {
	r, err := client.CreateNewUser(*ctx, &pb.CreateUserParams{FullName: fullName, Username: userName, HashToken: hashToken})
	if err != nil {
		log.Printf("failed to create new user: %v", err)
		return nil, err
	}

	return r, nil
}

// NoDuplicateUsersTest ensures you cannot create users with duplicate usernames.

func TestUserMethods(ctx context.Context, cancel context.CancelFunc) error {
	AssertClientNotNil()
	token := utils.HashText("Password")
	print(token)

	// Test if a user is created
	user1 := &pb.User{FullName: time.DateTime, Username: "Test User 1"}
	_, err := createUserWithParams(&ctx, user1.FullName, user1.Username, token)
	if err != nil {
		log.Printf("TestUserMethods Failed: CreateUserWithParams failed: %v", err)
		return err
	}
	log.Print("Passed: User created successfully.")

	// Ensure attempts to create users with duplicate names fails
	_, err = createUserWithParams(&ctx, user1.FullName, user1.Username, token)
	if err == nil {
		log.Printf("TestUserMethods Failed: Duplicate users were allowed.")
		return err
	}
	log.Printf("Passed: No duplicate users are allowed.")

	// Ensure User was added to db
	user2, err := client.GetUser(ctx, &pb.GetUserParams{Username: user1.GetUsername(), HashToken: token})
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
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
