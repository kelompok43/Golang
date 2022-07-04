package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelompok43/Golang/admin"
	"github.com/kelompok43/Golang/auth"
	authMiddleware "github.com/kelompok43/Golang/auth/middlewares"
	"github.com/kelompok43/Golang/config"
	"github.com/kelompok43/Golang/membership"
	"github.com/kelompok43/Golang/payment_method"
	"github.com/kelompok43/Golang/trainer"
	"github.com/kelompok43/Golang/transaction"
	"github.com/kelompok43/Golang/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	membership := membership.NewMembershipFactory(db)
	transaction := transaction.NewTransactionFactory(db, configJWT)

	e := echo.New()
	authMiddleware.LogMiddlewares(e)
	cJWT := configJWT.Init()

	userGroup := e.Group("/user")
	userGroup.GET("", user.GetAllData)
	userGroup.GET("/:id", user.GetByID, middleware.JWTWithConfig(cJWT))
	userGroup.GET("/profile/:id", user.GetByID, middleware.JWTWithConfig(cJWT))
	userGroup.POST("/profile/detail/:id", user.AddDetail)
	userGroup.GET("/get-email", user.GetByEmail)
	userGroup.PUT("/change-password/:id", user.ChangePassword)
	userGroup.PUT("/membership/status/:id", user.UpdateStatus)
	userGroup.POST("/login", user.Login)
	userGroup.POST("/register", user.Register)

	adminGroup := e.Group("/admin")
	adminGroup.GET("", admin.GetAllData)
	adminGroup.GET("/:id", admin.GetByID)
	adminGroup.GET("/get-email", admin.GetByEmail)
	adminGroup.PUT("/change-password/:id", admin.ChangePassword)
	adminGroup.POST("/login", admin.Login)
	adminGroup.POST("/register", admin.Register)

	trainerGroup := e.Group("/trainer")
	trainerGroup.POST("", trainer.AddData)
	trainerGroup.GET("", trainer.GetAllData)
	trainerGroup.GET("/:id", trainer.GetByID)
	trainerGroup.PUT("/:id", trainer.UpdateData)
	trainerGroup.DELETE("/:id", trainer.DeleteData)

	paymentGroup := e.Group("/payment")
	paymentGroup.POST("/method", paymentMethod.AddData)
	paymentGroup.GET("/method", paymentMethod.GetAllData)
	paymentGroup.GET("/method/:id", paymentMethod.GetByID)
	paymentGroup.PUT("/method/:id", paymentMethod.UpdateData)
	paymentGroup.DELETE("/method/:id", paymentMethod.DeleteData)

	transactionGroup := e.Group("/transaction", middleware.JWTWithConfig(cJWT))
	transactionGroup.POST("", transaction.AddData)
	transactionGroup.GET("", transaction.GetAllData)
	transactionGroup.GET("/:id", transaction.GetByID)
	// transactionGroup.GET("", transaction.GetByID) // /transaction?user_id=1
	// transactionGroup.GET("", transaction.GetByID) // /transaction?user_id=1&?trx_id=1
	transactionGroup.PUT("/:id", transaction.UpdateStatus)
	transactionGroup.DELETE("/:id", transaction.DeleteData)

	membershipGroup := e.Group("/membership")
	membershipGroup.POST("/category", membership.AddCategory)
	membershipGroup.GET("/category", membership.GetAllCategory)
	membershipGroup.GET("/category/:id", membership.GetCategoryByID)
	membershipGroup.PUT("/category/:id", membership.UpdateCategory)
	membershipGroup.DELETE("/category/:id", membership.DeleteCategory)

	e.Start(":9700")
}
