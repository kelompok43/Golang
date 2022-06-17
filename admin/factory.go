package admin

import (
	handlerAPI "github.com/kelompok43/Golang/admin/handler/api"
	repoMySQL "github.com/kelompok43/Golang/admin/repository/mysql"
	service "github.com/kelompok43/Golang/admin/service"
	authMiddleware "github.com/kelompok43/Golang/auth/middlewares"
	"gorm.io/gorm"
)

func NewAdminFactory(db *gorm.DB, configJWT authMiddleware.ConfigJWT) (adminHandler handlerAPI.AdminHandler) {
	adminRepo := repoMySQL.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepo, configJWT)
	adminHandler = handlerAPI.NewAdminHandler(adminService)
	return
}
