package service

import (
	"golang.org/x/net/context"
	"log"

	"github.com/shaun-golang/micro-service/application"
	proto "github.com/shaun-golang/micro-service/interface/grpc/proto/generated"
)

// Server gRPC
type Server struct {
	us application.UserAppInterface
}

//Users constructor
func NewUsers(us application.UserAppInterface) *Server {
	return &Server{
		us: us,
	}
}

//GetUser from id
func (s *Server) GetUser(ctx context.Context, id *proto.Id) (*proto.User, error) {

	log.Printf("Received message body from client: %d", id.Id)

	user, err := s.us.GetUser(uint64(id.GetId()))
	if err != nil {
		return &proto.User{Id: 1, FirstName: "ZhouXiaoyang"}, nil
	}

	return &proto.User{Id: int64(user.ID), FirstName: user.FirstName, LastName: user.LastName}, nil
}
