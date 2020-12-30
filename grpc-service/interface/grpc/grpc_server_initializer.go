package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	proto "github.com/shaun-golang/micro-service/interface/grpc/proto/generated"
	"github.com/shaun-golang/micro-service/interface/grpc/service"
)

// InitGrpcServer from config
func InitGrpcServer() {
	listener, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatalf("Failed to listen on port 9090: %v", err)
	}

	grpcServer := grpc.NewServer()

	s := service.Server{}
	proto.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9090: %v", err)
	}
}
