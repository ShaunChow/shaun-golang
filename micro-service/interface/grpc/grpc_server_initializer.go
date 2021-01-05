package grpc

import (
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/shaun-golang/micro-service/domain"
	proto "github.com/shaun-golang/micro-service/interface/grpc/proto/generated"
	"github.com/shaun-golang/micro-service/interface/grpc/service"
)

// InitGrpcServer from config
func InitGrpcServer() {
	listener, err := net.Listen("tcp", ":"+viper.GetString("grpc.server.port"))

	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", viper.GetString("grpc.server.port"), err)
	}

	grpcServer := grpc.NewServer()

	s := service.NewUsers(domain.UserRepository)

	proto.RegisterUserServiceServer(grpcServer, s)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", viper.GetString("grpc.server.port"), err)
	}
}
