package main

import (
    "github.com/spf13/pflag"

	"github.com/shaun-golang/micro-service/domain"
	"github.com/shaun-golang/micro-service/infrastructure"
	"github.com/shaun-golang/micro-service/interface/rest"
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
	domain.InitRepository()
	rest.InitGin()
}
