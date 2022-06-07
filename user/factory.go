package user

import (
	handlerAPI "github.com/kelompok43/Golang/user/handler/api"
	repoMySQL "github.com/kelompok43/Golang/user/repository/mysql"
	service "github.com/kelompok43/Golang/user/service"
	"gorm.io/gorm"
)

func NewUserFactory(db *gorm.DB) (userHandler handlerAPI.UserHandler) {
	userRepo := repoMySQL.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler = handlerAPI.NewUserHandler(userService)
	return
}
