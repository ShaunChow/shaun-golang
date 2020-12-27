package infrastructure

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/spf13/viper"

	"github.com/shaun-golang/micro-service/domain/entity"
)

// DB database connection
var DB *gorm.DB

// InitDB from Config
func InitDB() {

	db, err := gorm.Open(sqlite.Open(viper.GetString("datasource.sqlite.filepath")), &gorm.Config{})
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
