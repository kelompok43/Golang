package repoMySQL

import (
	"github.com/kelompok43/Golang/book/domain"
	"gorm.io/gorm"
)

type BookOnlineClass struct {
	gorm.Model
	ID        int
	UserID    int
	ClassID   int
	CreatedAt string
	UpdatedAt string
}

type BookOfflineClass struct {
	gorm.Model
	ID        int
	UserID    int
	ClassID   int
	CreatedAt string
	UpdatedAt string
}

func toOnlineClassDomain(rec BookOnlineClass) domain.OnlineClass {
	return domain.OnlineClass{
		ID:        rec.ID,
		UserID:    rec.UserID,
		ClassID:   rec.ClassID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func toOfflineClassDomain(rec BookOfflineClass) domain.OfflineClass {
	return domain.OfflineClass{
		ID:        rec.ID,
		UserID:    rec.UserID,
		ClassID:   rec.ClassID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromOnlineClassDomain(rec domain.OnlineClass) BookOnlineClass {
	return BookOnlineClass{
		ID:        rec.ID,
		UserID:    rec.UserID,
		ClassID:   rec.ClassID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromOfflineClassDomain(rec domain.OfflineClass) BookOfflineClass {
	return BookOfflineClass{
		ID:        rec.ID,
		UserID:    rec.UserID,
		ClassID:   rec.ClassID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
