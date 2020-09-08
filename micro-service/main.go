package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/shaun-golang/micro-service/domain/entity"
	"github.com/shaun-golang/micro-service/infrastructure/persistence"
	"github.com/shaun-golang/micro-service/interface"
)

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepository := persistence.NewUserRepository(db)

	users := interfaces.NewUsers(userRepository)

	// Migrate the schema
	db.AutoMigrate(&entity.User{})

	server := gin.Default()

	//user routes
	server.POST("/users", users.SaveUser)
	server.GET("/users", users.GetUsers)
	server.GET("/users/:user_id", users.GetUser)

	server.Run(":8080")
}
