package user

import (
	authMiddleware "github.com/kelompok43/Golang/auth/middlewares"
	repoMySQLM "github.com/kelompok43/Golang/membership/repository/mysql"
	serviceM "github.com/kelompok43/Golang/membership/service"
	handlerAPI "github.com/kelompok43/Golang/user/handler/api"
	repoMySQL "github.com/kelompok43/Golang/user/repository/mysql"
	service "github.com/kelompok43/Golang/user/service"
	"gorm.io/gorm"
)

func NewUserFactory(db *gorm.DB, configJWT authMiddleware.ConfigJWT) (userHandler handlerAPI.UserHandler) {
	membershipRepo := repoMySQLM.NewMembershipRepository(db)
	membershipService := serviceM.NewMembershipService(membershipRepo)
	userRepo := repoMySQL.NewUserRepository(db)
	userService := service.NewUserService(userRepo, configJWT, membershipService)
	userHandler = handlerAPI.NewUserHandler(userService)
	return
}
