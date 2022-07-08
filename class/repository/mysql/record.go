package repoMySQL

import (
	"github.com/kelompok43/Golang/class/domain"
	"gorm.io/gorm"
)

type ClassCategory struct {
	gorm.Model
	ID             int
	Name           string
	Description    string
	PictureLink    string
	CreatedAt      string
	UpdatedAt      string
	OnlineClasses  []OnlineClass  `gorm:"foreignKey:ClassCategoryID"`
	OfflineClasses []OfflineClass `gorm:"foreignKey:ClassCategoryID"`
}

type OnlineClass struct {
	gorm.Model
	ID              int
	ClassCategoryID int
	TrainerID       int
	Date            string
	StartedAt       string
	EndedAt         string
	Link            string
	CreatedAt       string
	UpdatedAt       string
}

type OfflineClass struct {
	gorm.Model
	ID              int
	ClassCategoryID int
	TrainerID       int
	Date            string
	StartedAt       string
	EndedAt         string
	Place           string
	Quota           int
	CreatedAt       string
	UpdatedAt       string
}

func toCategoryDomain(rec ClassCategory) domain.Category {
	return domain.Category{
		ID:          rec.ID,
		Name:        rec.Name,
		Description: rec.Description,
		PictureLink: rec.PictureLink,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func toOnlineDomain(rec OnlineClass) domain.Online {
	return domain.Online{
		ID:              rec.ID,
		ClassCategoryID: rec.ClassCategoryID,
		TrainerID:       rec.TrainerID,
		Date:            rec.Date,
		StartedAt:       rec.StartedAt,
		EndedAt:         rec.EndedAt,
		Link:            rec.Link,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
	}
}

func toOfflineDomain(rec OfflineClass) domain.Offline {
	return domain.Offline{
		ID:              rec.ID,
		ClassCategoryID: rec.ClassCategoryID,
		TrainerID:       rec.TrainerID,
		Date:            rec.Date,
		StartedAt:       rec.StartedAt,
		EndedAt:         rec.EndedAt,
		Place:           rec.Place,
		Quota:           rec.Quota,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
	}
}

func fromCategoryDomain(rec domain.Category) ClassCategory {
	return ClassCategory{
		ID:          rec.ID,
		Name:        rec.Name,
		Description: rec.Description,
		PictureLink: rec.PictureLink,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func fromOnlineDomain(rec domain.Online) OnlineClass {
	return OnlineClass{
		ID:              rec.ID,
		ClassCategoryID: rec.ClassCategoryID,
		TrainerID:       rec.TrainerID,
		Date:            rec.Date,
		StartedAt:       rec.StartedAt,
		EndedAt:         rec.EndedAt,
		Link:            rec.Link,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
	}
}

func fromOfflineDomain(rec domain.Offline) OfflineClass {
	return OfflineClass{
		ID:              rec.ID,
		ClassCategoryID: rec.ClassCategoryID,
		TrainerID:       rec.TrainerID,
		Date:            rec.Date,
		StartedAt:       rec.StartedAt,
		EndedAt:         rec.EndedAt,
		Place:           rec.Place,
		Quota:           rec.Quota,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
	}
}
