package chat_app_grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/jmagazine/chat-app-grpc/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	// Dial connection to grpc server
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewChatServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]string)
	new_users["Alice"] = "aliceloves123"
	new_users["Bob"] = "bobhates456"
	for fullname, username := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.CreateUserParams{FullName: fullname, Username: username, Password: ""})
		if err != nil {
			log.Fatalf("could not create new user: %v", err)
		}
		log.Printf(`
User Details:
Full Name: %s
Username: %s
Id: %d`, r.GetFullName(), r.GetUsername(), r.GetId())

	}
	params := &pb.GetAllUsersParams{}
	r, err := c.GetAllUsers(ctx, params)
	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}
	log.Print("\nUSERLIST: \n")
	fmt.Printf("r.GetUsers(): %v\n", r.GetUsers())
}
