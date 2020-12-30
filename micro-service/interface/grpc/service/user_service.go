package service

import (
	"log"

	"golang.org/x/net/context"

	proto "github.com/shaun-golang/micro-service/interface/grpc/proto/generated"
)

// Server gRPC
type Server struct {
}

//GetUser from id
func (s *Server) GetUser(ctx context.Context, id *proto.Id) (*proto.User, error) {

	log.Printf("Received message body from client: %d", id.Id)

	return &proto.User{Id: 1, FirstName: "ZhouXiaoyang"}, nil
}
