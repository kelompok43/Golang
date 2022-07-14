package handlerAPI

import (
	"time"

	helperTime "github.com/kelompok43/Golang/helpers/time"
	"github.com/kelompok43/Golang/membership/domain"
)

type RequestCategoryJSON struct {
	Category string `json:"category" form:"category" validate:"required"`
	Price    int    `json:"price" form:"price" validate:"required"`
	Duration int    `json:"duration" form:"duration" validate:"required"`
}

func toCategoryDomain(req RequestCategoryJSON) domain.MembershipCategory {
	return domain.MembershipCategory{
		Category: req.Category,
		Price:    req.Price,
		Duration: req.Duration,
	}
}

type ResponseCategoryJSON struct {
	Id        int       `json:"id"`
	Category  string    `json:"category" form:"category"`
	Price     int       `json:"price" form:"price"`
	Duration  int       `json:"duration" form:"duration"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type ResponseJSON struct {
	Id                   int       `json:"id"`
	UserID               int       `json:"user_id"`
	MembershipCategoryID int       `json:"membership_category_id"`
	ExpiredAt            time.Time `json:"expired_at"`
	CreatedAt            time.Time `json:"created_at" form:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" form:"updated_at"`
}

func fromCategoryDomain(domain domain.MembershipCategory) ResponseCategoryJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseCategoryJSON{
		Id:        domain.ID,
		Category:  domain.Category,
		Price:     domain.Price,
		Duration:  domain.Duration,
		CreatedAt: tmCreatedAt,
		UpdatedAt: tmUpdatedAt,
	}
}

func fromDomain(domain domain.Membership) ResponseJSON {
	tmExpiredAt := helperTime.NanoToTime(domain.ExpiredAt)
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		Id:                   domain.ID,
		UserID:               domain.UserID,
		MembershipCategoryID: domain.MembershipCategoryID,
		ExpiredAt:            tmExpiredAt,
		CreatedAt:            tmCreatedAt,
		UpdatedAt:            tmUpdatedAt,
	}
}
