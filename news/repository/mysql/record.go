package repoMySQL

import (
	"github.com/kelompok43/Golang/news/domain"
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	ID             int
	NewsCategoryID int
	Title          string
	Author         string
	Date           string
	Description    string
	PictureLink    string
	CreatedAt      string
	UpdatedAt      string
}

type NewsCategory struct {
	gorm.Model
	ID        int
	Name      string
	CreatedAt string
	UpdatedAt string
	News      News
}

func toDomain(rec News) domain.News {
	return domain.News{
		ID:             rec.ID,
		NewsCategoryID: rec.NewsCategoryID,
		Title:          rec.Title,
		Author:         rec.Author,
		Date:           rec.Date,
		Description:    rec.Description,
		PictureLink:    rec.PictureLink,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
	}
}

func toCategoryDomain(rec NewsCategory) domain.Category {
	return domain.Category{
		ID:        rec.ID,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(rec domain.News) News {
	return News{
		ID:             rec.ID,
		NewsCategoryID: rec.NewsCategoryID,
		Title:          rec.Title,
		Author:         rec.Author,
		Date:           rec.Date,
		Description:    rec.Description,
		PictureLink:    rec.PictureLink,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
	}
}

func fromCategoryDomain(rec domain.Category) NewsCategory {
	return NewsCategory{
		ID:        rec.ID,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
