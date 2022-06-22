package handlerAPI

import (
	"time"

	"github.com/kelompok43/Golang/admin/domain"
	helperTime "github.com/kelompok43/Golang/helpers/time"
)

type RequestJSON struct {
	Name     string `json:"name" form:"name" validate:"required"`
	DOB      string `json:"dob" form:"dob" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Role     string `json:"role" form:"role" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type RequestLoginJSON struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type Token struct {
	Token string `json:"token"`
}

func toDomain(req RequestJSON) domain.Admin {
	return domain.Admin{
		Name:     req.Name,
		DOB:      req.DOB,
		Gender:   req.Gender,
		Address:  req.Address,
		Role:     req.Role,
		Email:    req.Email,
		Password: req.Password,
	}
}

type ResponseJSON struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" form:"name"`
	DOB       string    `json:"dob" form:"dob"`
	Gender    string    `json:"gender" form:"gender"`
	Address   string    `json:"address" form:"address"`
	Role      string    `json:"role" form:"role"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.Admin) ResponseJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		Id:        domain.ID,
		Name:      domain.Name,
		DOB:       domain.DOB,
		Gender:    domain.Gender,
		Address:   domain.Address,
		Role:      domain.Role,
		Email:     domain.Email,
		CreatedAt: tmCreatedAt,
		UpdatedAt: tmUpdatedAt,
	}
}
