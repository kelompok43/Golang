package book

import (
	handlerAPI "github.com/kelompok43/Golang/book/handler/api"
	repoMySQL "github.com/kelompok43/Golang/book/repository/mysql"
	service "github.com/kelompok43/Golang/book/service"
	"gorm.io/gorm"
)

func NewBookFactory(db *gorm.DB) (bookHandler handlerAPI.BookHandler) {
	bookRepo := repoMySQL.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler = handlerAPI.NewBookHandler(bookService)
	return
}
