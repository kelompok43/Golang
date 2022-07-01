package handlerAPI

import (
	"time"

	helperTime "github.com/kelompok43/Golang/helpers/time"
	"github.com/kelompok43/Golang/payment_method/domain"
)

type RequestJSON struct {
	Name      string `json:"name" form:"name" validate:"required"`
	AccNumber string `json:"acc_number" form:"acc_number" validate:"required"`
	AccName   string `json:"acc_name" form:"acc_name" validate:"required"`
}

func toDomain(req RequestJSON) domain.PaymentMethod {
	return domain.PaymentMethod{
		Name:      req.Name,
		AccNumber: req.AccNumber,
		AccName:   req.AccName,
	}
}

type ResponseJSON struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" form:"name"`
	AccNumber string    `json:"acc_number" form:"acc_number"`
	AccName   string    `json:"acc_name" form:"acc_name"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.PaymentMethod) ResponseJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		Id:        domain.ID,
		Name:      domain.Name,
		AccNumber: domain.AccNumber,
		AccName:   domain.AccName,
		CreatedAt: tmCreatedAt,
		UpdatedAt: tmUpdatedAt,
	}
}
