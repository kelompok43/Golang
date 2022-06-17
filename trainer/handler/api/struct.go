package handlerAPI

import (
	"time"

	helperTime "github.com/kelompok43/Golang/helpers/time"
	"github.com/kelompok43/Golang/trainer/domain"
)

type RequestJSON struct {
	Name    string `json:"name" form:"name" validate:"required"`
	DOB     string `json:"dob" form:"dob" validate:"required"`
	Email   string `json:"email" form:"email" validate:"required,email"`
	Gender  string `json:"gender" form:"gender" validate:"required"`
	Phone   string `json:"phone" form:"phone" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	Picture string `json:"picture" form:"picture" validate:"required"`
	Field   string `json:"field" form:"field" validate:"required"`
}

func toDomain(req RequestJSON) domain.Trainer {
	return domain.Trainer{
		Name:    req.Name,
		DOB:     req.DOB,
		Email:   req.Email,
		Gender:  req.Gender,
		Address: req.Address,
		Phone:   req.Phone,
		Picture: req.Picture,
		Field:   req.Field,
	}
}

type ResponseJSON struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" form:"name"`
	DOB       string    `json:"dob" form:"dob"`
	Email     string    `json:"email" form:"email"`
	Gender    string    `json:"gender" form:"gender"`
	Phone     string    `json:"phone" form:"phone"`
	Address   string    `json:"address" form:"address"`
	Picture   string    `json:"picture" form:"picture"`
	Field     string    `json:"field" form:"field"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.Trainer) ResponseJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		Id:        domain.ID,
		Name:      domain.Name,
		DOB:       domain.DOB,
		Email:     domain.Email,
		Gender:    domain.Gender,
		Phone:     domain.Phone,
		Address:   domain.Address,
		Picture:   domain.Picture,
		Field:     domain.Field,
		CreatedAt: tmCreatedAt,
		UpdatedAt: tmUpdatedAt,
	}
}
