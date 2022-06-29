package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelompok43/Golang/admin"
	"github.com/kelompok43/Golang/auth"
	authMiddleware "github.com/kelompok43/Golang/auth/middlewares"
	"github.com/kelompok43/Golang/config"
	"github.com/kelompok43/Golang/payment_method"
	"github.com/kelompok43/Golang/trainer"
	"github.com/kelompok43/Golang/transaction"
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

	configJWT := authMiddleware.ConfigJWT{
		SecretJWT:       auth.SECRET_KEY,
		ExpiresDuration: auth.EXPIRED,
	}

	user := user.NewUserFactory(db, configJWT)
	admin := admin.NewAdminFactory(db, configJWT)
	trainer := trainer.NewTrainerFactory(db)
	paymentMethod := payment_method.NewPaymentMethodFactory(db)
	transaction := transaction.NewTransactionFactory(db)

	e := echo.New()

	e.GET("/user", user.GetAllData)
	e.GET("/user/:id", user.GetByID)
	e.GET("/user/profile/:id", user.GetByID)
	e.POST("/user/profile/detail/:id", user.AddDetail)
	e.GET("/user/forgot-password", user.GetByEmail)
	e.PUT("/user/change-password/:id", user.ChangePassword)
	e.POST("/user/login", user.Login)
	e.POST("/user/register", user.Register)

	e.GET("/admin/", admin.GetAllData)
	e.GET("/admin/:id", admin.GetByID)
	e.GET("/admin/forgot-password", admin.GetByEmail)
	e.PUT("/admin/change-password/:id", admin.ChangePassword)
	e.POST("/admin/login", admin.Login)
	e.POST("/admin/register", admin.Register)

	e.POST("/trainer", trainer.AddData)
	e.GET("/trainer", trainer.GetAllData)
	e.GET("/trainer/:id", trainer.GetByID)
	e.PUT("/trainer/:id", trainer.UpdateData)
	e.DELETE("/trainer/:id", trainer.DeleteData)

	e.POST("/payment/method", paymentMethod.AddData)
	e.GET("/payment/method", paymentMethod.GetAllData)
	e.GET("/payment/method/:id", paymentMethod.GetByID)
	e.PUT("/payment/method/:id", paymentMethod.UpdateData)
	e.DELETE("/payment/method/:id", paymentMethod.DeleteData)

	e.POST("/transaction", transaction.AddData)
	e.GET("/transaction", transaction.GetAllData)
	e.GET("/transaction/:id", transaction.GetByID)
	e.PUT("/transaction/:id", transaction.UpdateData)
	e.DELETE("/transaction/:id", transaction.DeleteData)

	e.Start(":9700")
}
