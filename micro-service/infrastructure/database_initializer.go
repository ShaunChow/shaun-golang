package infrastructure

import (
	"github.com/shaun-golang/micro-service/domain/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	db, err := gorm.Open(sqlite.Open("./resource/database/sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrate the schema
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic("Failed to migrate the schema to database!")
	}

	DB = db
}
