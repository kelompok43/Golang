package payment_method

import (
	handlerAPI "github.com/kelompok43/Golang/payment_method/handler/api"
	repoMySQL "github.com/kelompok43/Golang/payment_method/repository/mysql"
	service "github.com/kelompok43/Golang/payment_method/service"
	"gorm.io/gorm"
)

func NewPaymentMethodFactory(db *gorm.DB) (paymentMethodHandler handlerAPI.PaymentMethodHandler) {
	paymentMethodRepo := repoMySQL.NewPaymentMethodRepository(db)
	paymentMethodService := service.NewPaymentMethodService(paymentMethodRepo)
	paymentMethodHandler = handlerAPI.NewPaymentMethodHandler(paymentMethodService)
	return
}
