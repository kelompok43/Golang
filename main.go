package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelompok43/Golang/admin"
	"github.com/kelompok43/Golang/auth"
	authMiddleware "github.com/kelompok43/Golang/auth/middlewares"
	"github.com/kelompok43/Golang/book"
	"github.com/kelompok43/Golang/class"
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
	class := class.NewClassFactory(db)
	book := book.NewBookFactory(db)

	e := echo.New()
	e.Use(middleware.CORS())
	authMiddleware.LogMiddlewares(e)
	cJWT := configJWT.Init()

	// member := os.Getenv("USER_STATUS_MEMBER")
	// notMember := os.Getenv("USER_STATUS_NOT_MEMBER")
	// superAdmin := os.Getenv("SUPER_ADMIN")
	// operationalAdmin := os.Getenv("OPERATIONAL_ADMIN")

	userGroup := e.Group("/user")
	userGroup.GET("", user.GetAllData, middleware.JWTWithConfig(cJWT))
	userGroup.GET("/:id", user.GetByID, middleware.JWTWithConfig(cJWT))
	// userGroup.GET("/profile/:id", user.GetByID, middleware.JWTWithConfig(cJWT))
	userGroup.POST("/detail/:id", user.AddDetail, middleware.JWTWithConfig(cJWT))
	userGroup.PUT("/detail/:id", user.Update, middleware.JWTWithConfig(cJWT))
	userGroup.GET("/:id/transaction", transaction.GetUserTrx)
	userGroup.GET("/:id/transaction/:trx_id", transaction.GetUserTrxByID)
	userGroup.GET("/get-email", user.GetByEmail)
	userGroup.PUT("/:id/change-password", user.ChangePassword)
	userGroup.GET("/:id/membership", membership.GetByUserID)
	userGroup.PUT("/:id/membership/status", user.UpdateStatus)
	userGroup.POST("/login", user.Login)
	userGroup.POST("/register", user.Register)

	adminGroup := e.Group("/admin")
	adminGroup.GET("", admin.GetAllData, middleware.JWTWithConfig(cJWT))
	adminGroup.GET("/:id", admin.GetByID, middleware.JWTWithConfig(cJWT))
	adminGroup.GET("/get-email", admin.GetByEmail)
	adminGroup.PUT("/change-password/:id", admin.ChangePassword)
	adminGroup.POST("/login", admin.Login)
	adminGroup.POST("/register", admin.Register)

	trainerGroup := e.Group("/trainer", middleware.JWTWithConfig(cJWT))
	trainerGroup.POST("", trainer.AddData)
	trainerGroup.GET("", trainer.GetAllData)
	trainerGroup.GET("/:id", trainer.GetByID)
	trainerGroup.PUT("/:id", trainer.UpdateData)
	trainerGroup.DELETE("/:id", trainer.DeleteData)

	paymentGroup := e.Group("/payment", middleware.JWTWithConfig(cJWT))
	paymentGroup.POST("/method", paymentMethod.AddData)
	paymentGroup.GET("/method", paymentMethod.GetAllData)
	paymentGroup.GET("/method/:id", paymentMethod.GetByID)
	paymentGroup.PUT("/method/:id", paymentMethod.UpdateData)
	paymentGroup.DELETE("/method/:id", paymentMethod.DeleteData)

	transactionGroup := e.Group("/transaction", middleware.JWTWithConfig(cJWT))
	transactionGroup.POST("", transaction.AddData)
	transactionGroup.GET("", transaction.GetAllData)
	transactionGroup.GET("/:id", transaction.GetByID)
	transactionGroup.PUT("/:id", transaction.UpdateStatus)
	transactionGroup.DELETE("/:id", transaction.DeleteData)

	membershipGroup := e.Group("/membership")
	membershipGroup.GET("", membership.GetAllData)
	membershipGroup.GET("/:id", membership.GetByID)
	membershipGroup.POST("/category", membership.AddCategory, middleware.JWTWithConfig(cJWT))
	membershipGroup.GET("/category", membership.GetAllCategory)
	membershipGroup.GET("/category/:id", membership.GetCategoryByID)
	membershipGroup.PUT("/category/:id", membership.UpdateCategory, middleware.JWTWithConfig(cJWT))
	membershipGroup.DELETE("/category/:id", membership.DeleteCategory, middleware.JWTWithConfig(cJWT))

	classGroup := e.Group("/class")
	classGroup.POST("/category", class.AddCategory)
	classGroup.GET("/category", class.GetAllCategory)
	classGroup.GET("/category/:id", class.GetCategoryByID)
	classGroup.PUT("/category/:id", class.UpdateCategory)
	classGroup.DELETE("/category/:id", class.DeleteCategory)
	classGroup.POST("/online", class.AddOnline)
	classGroup.GET("/online", class.GetAllOnline)
	classGroup.GET("/online/:id", class.GetOnlineByID)
	classGroup.PUT("/online/:id", class.UpdateOnline)
	classGroup.DELETE("/online/:id", class.DeleteOnline)
	classGroup.POST("/offline", class.AddOffline)
	classGroup.GET("/offline", class.GetAllOffline)
	classGroup.GET("/offline/:id", class.GetOfflineByID)
	classGroup.PUT("/offline/:id", class.UpdateOffline)
	classGroup.DELETE("/offline/:id", class.DeleteOffline)

	bookGroup := e.Group("/book")
	bookGroup.POST("/online-class", book.AddOnlineClass)
	bookGroup.GET("/online-class", book.GetAllOnlineClass)
	bookGroup.GET("/online-class/:id", book.GetOnlineClassByID)
	bookGroup.PUT("/online-class/:id", book.UpdateOnlineClass)
	bookGroup.DELETE("/online-class/:id", book.DeleteOnlineClass)
	bookGroup.POST("/offline-class", book.AddOfflineClass)
	bookGroup.GET("/offline-class", book.GetAllOfflineClass)
	bookGroup.GET("/offline-class/:id", book.GetOfflineClassByID)
	bookGroup.PUT("/offline-class/:id", book.UpdateOfflineClass)
	bookGroup.DELETE("/offline-class/:id", book.DeleteOfflineClass)

	e.Start(":9700")
}
