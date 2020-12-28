package infrastructure

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

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

var RedisClient *redis.Client

func RedisInit() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: viper.GetString("redis.auth"),
		DB:       viper.GetInt("redis.db"),
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}
}
