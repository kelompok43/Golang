package handlerAPI

import (
	"mime/multipart"
	"time"

	helperTime "github.com/kelompok43/Golang/helpers/time"
	"github.com/kelompok43/Golang/transaction/domain"
)

type RequestJSON struct {
	UserID          int `json:"user_id" form:"user_id" validate:"required"`
	PaymentMethodID int `json:"payment_method_id" form:"payment_method_id" validate:"required"`
	TotalPrice      int `json:"total_price" form:"total_price" validate:"required"`
	// Status          string         `json:"status" form:"status" validate:"required"`
	PaymentReceipt multipart.File `form:"payment_receipt"`
}

type RequestStatus struct {
	// TransactionID int
	Status string `json:"status" validate:"required"`
}

func toDomain(req RequestJSON) domain.Transaction {
	return domain.Transaction{
		UserID:          req.UserID,
		PaymentMethodID: req.PaymentMethodID,
		TotalPrice:      req.TotalPrice,
		// Status:          req.Status,
		Payment_Receipt: req.PaymentReceipt,
	}
}

func statusToDomain(req RequestStatus) domain.Transaction {
	return domain.Transaction{
		Status: req.Status,
	}
}

type ResponseJSON struct {
	Id              int       `json:"id"`
	UserID          int       `json:"user_id" form:"user_id"`
	PaymentMethodID int       `json:"payment_method_id" form:"payment_method_id"`
	TotalPrice      int       `json:"total_price" form:"total_price"`
	Status          string    `json:"status" form:"status"`
	PictureLink     string    `json:"payment_receipt" form:"payment_receipt"`
	CreatedAt       time.Time `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.Transaction) ResponseJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		Id:              domain.ID,
		UserID:          domain.UserID,
		PaymentMethodID: domain.PaymentMethodID,
		TotalPrice:      domain.TotalPrice,
		Status:          domain.Status,
		PictureLink:     domain.PictureLink,
		CreatedAt:       tmCreatedAt,
		UpdatedAt:       tmUpdatedAt,
	}
}

//response kompleks
// type ResponseJSON struct {
// 	Id              int           `json:"id"`
// 	UserID          int           `json:"user_id" form:"user_id"`
// 	User            User          `json:"user"`
// 	PaymentMethodID int           `json:"payment_method_id" form:"payment_method_id"`
// 	PaymentMethod   PaymentMethod `json:"payment_method"`
// 	TotalPrice      int           `json:"total_price" form:"total_price"`
// 	Status          string        `json:"status" form:"status"`
// 	PictureLink     string        `json:"payment_receipt" form:"payment_receipt"`
// 	CreatedAt       time.Time     `json:"created_at" form:"created_at"`
// 	UpdatedAt       time.Time     `json:"updated_at" form:"updated_at"`
// }

// type User struct {
// 	Id        int       `json:"id"`
// 	Name      string    `json:"name" form:"name"`
// 	DOB       string    `json:"dob" form:"dob"`
// 	Email     string    `json:"email" form:"email"`
// 	Phone     string    `json:"phone" form:"phone"`
// 	Address   string    `json:"address" form:"address"`
// 	Gender    string    `json:"gender" form:"gender"`
// 	Status    string    `json:"status" form:"status"`
// 	CreatedAt time.Time `json:"created_at" form:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
// }

// type PaymentMethod struct {
// 	Id        int       `json:"id"`
// 	Name      string    `json:"name" form:"name"`
// 	AccNumber string    `json:"acc_number" form:"acc_number"`
// 	AccName   string    `json:"acc_name" form:"acc_name"`
// 	CreatedAt time.Time `json:"created_at" form:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
// }

// func fromDomain(domain domain.Transaction, domainUser domainUser.User, domainPM domainPaymentMethod.PaymentMethod) ResponseJSON {
// 	//parse unix timestamp to time.Time
// 	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
// 	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

// 	return ResponseJSON{
// 		Id:     domain.ID,
// 		UserID: domain.UserID,
// 		User: User{
// 			Id:        domainUser.ID,
// 			Name:      domainUser.Name,
// 			DOB:       domainUser.DOB,
// 			Email:     domainUser.Email,
// 			Phone:     domainUser.Phone,
// 			Address:   domainUser.Address,
// 			Gender:    domainUser.Gender,
// 			Status:    domainUser.Status,
// 			CreatedAt: helperTime.NanoToTime(domainUser.CreatedAt),
// 			UpdatedAt: helperTime.NanoToTime(domainUser.UpdatedAt),
// 		},
// 		PaymentMethodID: domain.PaymentMethodID,
// 		TotalPrice:      domain.TotalPrice,
// 		Status:          domain.Status,
// 		PictureLink:     domain.PictureLink,
// 		CreatedAt:       tmCreatedAt,
// 		UpdatedAt:       tmUpdatedAt,
// 	}
// }
