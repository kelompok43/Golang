package news

import (
	handlerAPI "github.com/kelompok43/Golang/news/handler/api"
	repoMySQL "github.com/kelompok43/Golang/news/repository/mysql"
	service "github.com/kelompok43/Golang/news/service"
	"gorm.io/gorm"
)

func NewNewsFactory(db *gorm.DB) (newsHandler handlerAPI.NewsHandler) {
	newsRepo := repoMySQL.NewNewsRepository(db)
	newsService := service.NewNewsService(newsRepo)
	newsHandler = handlerAPI.NewNewsHandler(newsService)
	return
}
