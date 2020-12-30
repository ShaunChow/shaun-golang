package main

import (
	"github.com/shaun-golang/micro-service/domain"
	"github.com/shaun-golang/micro-service/infrastructure"
	"github.com/shaun-golang/micro-service/interface/grpc"
	"github.com/shaun-golang/micro-service/interface/rest"
	"github.com/spf13/pflag"
)

var (
	conf = pflag.StringP("config", "c", "", "config filepath")
)

func main() {

	// Init Config File
	pflag.Parse()
	if err := infrastructure.InitConfig(*conf); err != nil {
		panic(err)
	}

	infrastructure.InitDB()
	//infrastructure.RedisInit()

	domain.InitRepository()

	grpcClientSyn := make(chan bool, 1)
	go grpc.InitGrpcClient()

	grpcServerSyn := make(chan bool, 1)
	go grpc.InitGrpcServer()

	rest.InitGin()

	<-grpcServerSyn
	<-grpcClientSyn

}
