package transaction

import (
	authMiddleware "github.com/kelompok43/Golang/auth/middlewares"
	repoMySQLM "github.com/kelompok43/Golang/membership/repository/mysql"
	serviceM "github.com/kelompok43/Golang/membership/service"
	handlerAPI "github.com/kelompok43/Golang/transaction/handler/api"
	repoMySQL "github.com/kelompok43/Golang/transaction/repository/mysql"
	service "github.com/kelompok43/Golang/transaction/service"
	repoMySQLU "github.com/kelompok43/Golang/user/repository/mysql"
	serviceU "github.com/kelompok43/Golang/user/service"
	"gorm.io/gorm"
)

func NewTransactionFactory(db *gorm.DB, configJWT authMiddleware.ConfigJWT) (transactionHandler handlerAPI.TransactionHandler) {
	// pmRepo := repoMySQLPM.NewPaymentMethodRepository(db)
	// pmService := servicePM.NewPaymentMethodService(pmRepo)
	membershipRepo := repoMySQLM.NewMembershipRepository(db)
	membershipService := serviceM.NewMembershipService(membershipRepo)
	userRepo := repoMySQLU.NewUserRepository(db)
	userService := serviceU.NewUserService(userRepo, configJWT, membershipService)
	transactionRepo := repoMySQL.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo, membershipService, userService)
	transactionHandler = handlerAPI.NewTransactionHandler(transactionService)
	return
}
