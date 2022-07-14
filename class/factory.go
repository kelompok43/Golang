package class

import (
	handlerAPI "github.com/kelompok43/Golang/class/handler/api"
	repoMySQL "github.com/kelompok43/Golang/class/repository/mysql"
	service "github.com/kelompok43/Golang/class/service"
	"gorm.io/gorm"
)

func NewClassFactory(db *gorm.DB) (classHandler handlerAPI.ClassHandler) {
	classRepo := repoMySQL.NewClassRepository(db)
	classService := service.NewClassService(classRepo)
	classHandler = handlerAPI.NewClassHandler(classService)
	return
}
