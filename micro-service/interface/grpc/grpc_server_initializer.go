package grpc

import (
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/shaun-golang/micro-service/domain"
	"github.com/shaun-golang/micro-service/infrastructure"
	proto "github.com/shaun-golang/micro-service/interface/grpc/proto/generated"
	"github.com/shaun-golang/micro-service/interface/grpc/service"
)

// InitGrpcServer from config
func InitGrpcServer() {

	log := infrastructure.Logger

	listener, err := net.Listen("tcp", ":"+viper.GetString("grpc.server.port"))

	if err != nil {
		log.Sugar().Errorf("Failed to listen on port %s: %v", viper.GetString("grpc.server.port"), err)
	}

	grpcServer := grpc.NewServer()

	s := service.NewUsers(domain.UserRepository)

	proto.RegisterUserServiceServer(grpcServer, s)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Sugar().Errorf("Failed to serve gRPC server over port %s: %v", viper.GetString("grpc.server.port"), err)
	}
}
