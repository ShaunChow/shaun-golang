package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/shaun-golang/micro-service/domain"
	"github.com/shaun-golang/micro-service/interface/rest/handler"
)

// SERVER ginEngine
var SERVER *gin.Engine

// InitGin from config
func InitGin() {

	users := handler.NewUsers(domain.UserRepository)

	server := gin.Default()

	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	server.Use(cors.Default())

	//user routes
	server.POST("/users", users.SaveUser)
	server.GET("/users", users.GetUsers)
	server.GET("/users/:user_id", users.GetUser)

	err := server.Run(":" + viper.GetString("server.port"))
	if err != nil {
		panic("Failed to run gin on port:" + viper.GetString("server.port") + "!")
	}
}
