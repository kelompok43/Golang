package handlerAPI

import (
	"mime/multipart"
	"time"

	helperTime "github.com/kelompok43/Golang/helpers/time"
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

type RequestPasswordJSON struct {
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type RequestDetailJSON struct {
	UserID  int
	Picture multipart.File `form:"picture"`
	Name    string         `json:"name" form:"name" validate:"required"`
	DOB     string         `json:"dob" form:"dob" validate:"required"`
	Address string         `json:"address" form:"address" validate:"required"`
	Email   string         `json:"email" form:"email" validate:"required,email"`
	Phone   string         `json:"phone" form:"phone" validate:"required"`
	// Gender  string `json:"gender" form:"gender" validate:"required"`
}

type Token struct {
	Token string `json:"token"`
}

func pwdToDomain(req RequestPasswordJSON) domain.User {
	return domain.User{
		Password: req.Password,
	}
}

func toDomain(req RequestJSON) domain.User {
	return domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func detailToDomain(req RequestDetailJSON) domain.User {
	return domain.User{
		ID:      req.UserID,
		Picture: req.Picture,
		Name:    req.Name,
		DOB:     req.DOB,
		Address: req.Address,
		Email:   req.Email,
		Phone:   req.Phone,
	}
}

type ResponseJSON struct {
	Id      int    `json:"id"`
	Name    string `json:"name" form:"name"`
	DOB     string `json:"dob" form:"dob"`
	Email   string `json:"email" form:"email"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
	// Gender      string    `json:"gender" form:"gender"`
	Status      string    `json:"status" form:"status"`
	PictureLink string    `json:"picture" form:"picture"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.User) ResponseJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		Id:          domain.ID,
		PictureLink: domain.PictureLink,
		Name:        domain.Name,
		DOB:         domain.DOB,
		Email:       domain.Email,
		Phone:       domain.Phone,
		Address:     domain.Address,
		// Gender:    domain.Gender,
		Status:    domain.Status,
		CreatedAt: tmCreatedAt,
		UpdatedAt: tmUpdatedAt,
	}
}
