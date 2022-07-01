package trainer

import (
	handlerAPI "github.com/kelompok43/Golang/trainer/handler/api"
	repoMySQL "github.com/kelompok43/Golang/trainer/repository/mysql"
	service "github.com/kelompok43/Golang/trainer/service"
	"gorm.io/gorm"
)

func NewTrainerFactory(db *gorm.DB) (trainerHandler handlerAPI.TrainerHandler) {
	trainerRepo := repoMySQL.NewTrainerRepository(db)
	trainerService := service.NewTrainerService(trainerRepo)
	trainerHandler = handlerAPI.NewTrainerHandler(trainerService)
	return
}
