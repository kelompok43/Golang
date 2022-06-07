package handlerAPI

import (
	"time"

	"github.com/kelompok43/Golang/user/domain"
)

type RequestJSON struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type RequestLoginJSON struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

func toDomain(req RequestJSON) domain.User {
	return domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

type ResponseJSON struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" form:"name"`
	DOB       string    `json:"dob" form:"dob"`
	Email     string    `json:"email" form:"email"`
	Phone     string    `json:"phone" form:"phone"`
	Address   string    `json:"address" form:"address"`
	Gender    string    `json:"gender" form:"gender"`
	Status    string    `json:"status" form:"status"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.User) ResponseJSON {
	return ResponseJSON{
		Id:        domain.ID,
		Name:      domain.Name,
		DOB:       domain.DOB,
		Email:     domain.Email,
		Phone:     domain.Phone,
		Address:   domain.Address,
		Gender:    domain.Gender,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
