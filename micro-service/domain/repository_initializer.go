package domain

import (
	"github.com/shaun-golang/micro-service/infrastructure"
	"github.com/shaun-golang/micro-service/infrastructure/persistence"
)

var UserRepository *persistence.UserRepo

func InitRepository() {
	UserRepository = persistence.NewUserRepository(infrastructure.DB)
}
