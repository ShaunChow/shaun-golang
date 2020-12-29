package grpc

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	proto "github.com/shaun-golang/micro-service/interface/grpc/proto/generated"
)

// InitGrpcClient from config
func InitGrpcClient() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9090", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	c := proto.NewUserServiceClient(conn)

	id := proto.Id{
		Id: 1,
	}

	response, err := c.GetUser(context.Background(), &id)

	if err != nil {
		log.Fatalf("Error when calling GetUser: %s", err)

	}

	log.Printf("Response from Server: %s", response.FirstName)
}
