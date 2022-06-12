package user

import (
	authMiddleware "github.com/kelompok43/Golang/auth/middlewares"
	handlerAPI "github.com/kelompok43/Golang/user/handler/api"
	repoMySQL "github.com/kelompok43/Golang/user/repository/mysql"
	service "github.com/kelompok43/Golang/user/service"
	"gorm.io/gorm"
)

func NewUserFactory(db *gorm.DB, configJWT authMiddleware.ConfigJWT) (userHandler handlerAPI.UserHandler) {
	userRepo := repoMySQL.NewUserRepository(db)
	userService := service.NewUserService(userRepo, configJWT)
	userHandler = handlerAPI.NewUserHandler(userService)
	return
}
