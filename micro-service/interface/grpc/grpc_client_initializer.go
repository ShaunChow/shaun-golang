package grpc

import (
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/shaun-golang/micro-service/infrastructure"
	proto "github.com/shaun-golang/micro-service/interface/grpc/proto/generated"
)

// InitGrpcClient from config
func InitGrpcClient() {
	var conn *grpc.ClientConn

	log := infrastructure.Logger

	conn, err := grpc.Dial(viper.GetString("grpc.client.ip")+":"+viper.GetString("grpc.client.port"), grpc.WithInsecure())

	if err != nil {
		log.Sugar().Fatalf("Could not connect: %s", err)
	}

	defer conn.Close()

	c := proto.NewUserServiceClient(conn)

	id := proto.Id{
		Id: 2,
	}

	response, err := c.GetUser(context.Background(), &id)

	if err != nil {
		log.Sugar().Fatalf("Error when calling GetUser: %s", err)

	}

	log.Sugar().Info("Response from Server: %s", response.FirstName)
}
