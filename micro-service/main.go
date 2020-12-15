package main

import (
	"github.com/shaun-golang/micro-service/domain"
	"github.com/shaun-golang/micro-service/infrastructure"
	"github.com/shaun-golang/micro-service/interface/rest"
)

func main() {
	infrastructure.InitDB()
	domain.InitRepository()
	rest.InitGin()
}
