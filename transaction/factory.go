package transaction

import (
	handlerAPI "github.com/kelompok43/Golang/transaction/handler/api"
	repoMySQL "github.com/kelompok43/Golang/transaction/repository/mysql"
	service "github.com/kelompok43/Golang/transaction/service"
	"gorm.io/gorm"
)

func NewTransactionFactory(db *gorm.DB) (transactionHandler handlerAPI.TransactionHandler) {
	// userRepo := repoMySQLU.NewUserRepository(db)
	// userService := serviceU.NewUserService(userRepo, configJWT)
	// pmRepo := repoMySQLPM.NewPaymentMethodRepository(db)
	// pmService := servicePM.NewPaymentMethodService(pmRepo)
	transactionRepo := repoMySQL.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler = handlerAPI.NewTransactionHandler(transactionService)
	return
}
