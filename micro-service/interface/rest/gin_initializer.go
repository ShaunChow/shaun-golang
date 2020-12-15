package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/shaun-golang/micro-service/domain"
	"github.com/shaun-golang/micro-service/interface/rest/handler"
)

var SERVER *gin.Engine

func InitGin() {

	users := handler.NewUsers(domain.UserRepository)

	server := gin.Default()
	//user routes
	server.POST("/users", users.SaveUser)
	server.GET("/users", users.GetUsers)
	server.GET("/users/:user_id", users.GetUser)

	err := server.Run(":8080")
	if err != nil {
		panic("Failed to run gin on port:8080!")
	}
}
