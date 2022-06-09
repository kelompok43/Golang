package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelompok43/Golang/config"
	"github.com/kelompok43/Golang/user"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	config.Init()
	db := config.DBInit()
	config.DBMigrate(db)

	user := user.NewUserFactory(db)

	e := echo.New()

	e.GET("/user", user.GetAllData)
	e.GET("/user/:id", user.GetByID)
	e.POST("/user/register", user.Register)

	e.Start(":9700")
}
