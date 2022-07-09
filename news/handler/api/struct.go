package handlerAPI

import (
	"mime/multipart"
	"time"

	helperTime "github.com/kelompok43/Golang/helpers/time"
	"github.com/kelompok43/Golang/news/domain"
)

type RequestJSON struct {
	NewsCategoryID int            `json:"news_category_id" form:"news_category_id" validate:"required"`
	Title          string         `json:"title" form:"title" validate:"required"`
	Author         string         `json:"author" form:"author" validate:"required"`
	Date           string         `json:"date" form:"date" validate:"required"`
	Description    string         `json:"description" form:"description" validate:"required"`
	Picture        multipart.File `form:"picture"`
}

type RequestCategoryJSON struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func toDomain(req RequestJSON) domain.News {
	return domain.News{
		NewsCategoryID: req.NewsCategoryID,
		Title:          req.Title,
		Author:         req.Author,
		Date:           req.Date,
		Description:    req.Description,
		Picture:        req.Picture,
	}
}

func toCategoryDomain(req RequestCategoryJSON) domain.Category {
	return domain.Category{
		Name: req.Name,
	}
}

type ResponseJSON struct {
	Id             int       `json:"id"`
	NewsCategoryID int       `json:"news_category_id" form:"news_category_id"`
	Title          string    `json:"title" form:"title"`
	Author         string    `json:"author" form:"author"`
	Date           string    `json:"date" form:"date"`
	Description    string    `json:"description" form:"description"`
	PictureLink    string    `json:"picture" form:"picture"`
	CreatedAt      time.Time `json:"created_at" form:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" form:"updated_at"`
}

type ResponseCategoryJSON struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func fromDomain(domain domain.News) ResponseJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseJSON{
		Id:             domain.ID,
		NewsCategoryID: domain.NewsCategoryID,
		Title:          domain.Title,
		Author:         domain.Author,
		Date:           domain.Date,
		Description:    domain.Description,
		PictureLink:    domain.PictureLink,
		CreatedAt:      tmCreatedAt,
		UpdatedAt:      tmUpdatedAt,
	}
}

func fromCategoryDomain(domain domain.Category) ResponseCategoryJSON {
	//parse unix timestamp to time.Time
	tmCreatedAt := helperTime.NanoToTime(domain.CreatedAt)
	tmUpdatedAt := helperTime.NanoToTime(domain.UpdatedAt)

	return ResponseCategoryJSON{
		Id:        domain.ID,
		Name:      domain.Name,
		CreatedAt: tmCreatedAt,
		UpdatedAt: tmUpdatedAt,
	}
}
