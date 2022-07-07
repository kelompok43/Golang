package handlerAPI

import (
	"mime/multipart"
	"time"

	"github.com/kelompok43/Golang/class/domain"
	helperTime "github.com/kelompok43/Golang/helpers/time"
)

type RequestCategoryJSON struct {
	Name        string         `json:"name" form:"name" validate:"required"`
	Description string         `json:"description" form:"description" validate:"required"`
	Picture     multipart.File `form:"picture"`
}

type RequestOnlineJSON struct {
	ClassCategoryID int    `json:"class_category_id" form:"class_category_id" validate:"required"`
	TrainerID       int    `json:"trainer_id" form:"trainer_id" validate:"required"`
	Date            string `json:"date" form:"date" validate:"required"`
	StartedAt       string `json:"started_at" form:"started_at" validate:"required"`
	EndedAt         string `json:"ended_at" form:"ended_at" validate:"required"`
	Link            string `json:"link" form:"link" validate:"required"`
}

func toCategoryDomain(req RequestCategoryJSON) domain.Category {
	return domain.Category{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
	}
}

func toOnlineDomain(req RequestOnlineJSON) domain.Online {
	return domain.Online{
		ClassCategoryID: req.ClassCategoryID,
		TrainerID:       req.TrainerID,
		Date:            req.Date,
		StartedAt:       req.StartedAt,
		EndedAt:         req.EndedAt,
		Link:            req.Link,
	}
}

type ResponseCategoryJSON struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	PictureLink string    `json:"picture" form:"picture"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

type ResponseOnlineJSON struct {
	Id              int       `json:"id"`
	ClassCategoryID int       `json:"class_category_id" form:"class_category_id"`
	TrainerID       int       `json:"trainer_id" form:"trainer_id"`
	Date            time.Time `json:"date" form:"date"`
	StartedAt       time.Time `json:"started_at" form:"started_at"`
	EndedAt         time.Time `json:"ended_at" form:"ended_at"`
	Link            string    `json:"link" form:"link"`
	CreatedAt       time.Time `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" form:"updated_at"`
}

func fromCategoryDomain(domain domain.Category) ResponseCategoryJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseCategoryJSON{
		Id:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		PictureLink: domain.PictureLink,
		CreatedAt:   tmCreatedAt,
		UpdatedAt:   tmUpdatedAt,
	}
}

func fromOnlineDomain(domain domain.Online) ResponseOnlineJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)
	tmDate, _ := helperTime.JustDate(domain.Date)
	tmStartedAt, _ := helperTime.JustTime(domain.StartedAt)
	tmEndedAt, _ := helperTime.JustTime(domain.EndedAt)

	return ResponseOnlineJSON{
		Id:              domain.ID,
		ClassCategoryID: domain.ClassCategoryID,
		TrainerID:       domain.TrainerID,
		Date:            tmDate,
		StartedAt:       tmStartedAt,
		EndedAt:         tmEndedAt,
		Link:            domain.Link,
		CreatedAt:       tmCreatedAt,
		UpdatedAt:       tmUpdatedAt,
	}
}
