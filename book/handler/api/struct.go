package handlerAPI

import (
	"time"

	"github.com/kelompok43/Golang/book/domain"
	helperTime "github.com/kelompok43/Golang/helpers/time"
)

type RequestOnlineClassJSON struct {
	UserID  int `json:"user_id" form:"user_id" validate:"required"`
	ClassID int `json:"class_id" form:"class_id" validate:"required"`
}

type RequestOfflineClassJSON struct {
	UserID  int `json:"user_id" form:"user_id" validate:"required"`
	ClassID int `json:"class_id" form:"class_id" validate:"required"`
}

func toOnlineClassDomain(req RequestOnlineClassJSON) domain.OnlineClass {
	return domain.OnlineClass{
		UserID:        req.UserID,
		OnlineClassID: req.ClassID,
	}
}

func toOfflineClassDomain(req RequestOfflineClassJSON) domain.OfflineClass {
	return domain.OfflineClass{
		UserID:         req.UserID,
		OfflineClassID: req.ClassID,
	}
}

type ResponseOnlineClassJSON struct {
	Id        int       `json:"id"`
	UserID    int       `json:"user_id" form:"user_id"`
	ClassID   int       `json:"class_id" form:"class_id"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type ResponseOfflineClassJSON struct {
	Id        int       `json:"id"`
	UserID    int       `json:"user_id" form:"user_id"`
	ClassID   int       `json:"class_id" form:"class_id"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func fromOnlineClassDomain(domain domain.OnlineClass) ResponseOnlineClassJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseOnlineClassJSON{
		Id:        domain.ID,
		UserID:    domain.UserID,
		ClassID:   domain.OnlineClassID,
		CreatedAt: tmCreatedAt,
		UpdatedAt: tmUpdatedAt,
	}
}

func fromOfflineClassDomain(domain domain.OfflineClass) ResponseOfflineClassJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseOfflineClassJSON{
		Id:        domain.ID,
		UserID:    domain.UserID,
		ClassID:   domain.OfflineClassID,
		CreatedAt: tmCreatedAt,
		UpdatedAt: tmUpdatedAt,
	}
}
