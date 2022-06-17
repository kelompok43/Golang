package repoMySQL

import (
	"github.com/kelompok43/Golang/trainer/domain"
	"gorm.io/gorm"
)

type Trainer struct {
	gorm.Model
	ID        int
	Name      string
	Email     string
	DOB       string
	Gender    string
	Phone     string
	Address   string
	Picture   string
	Field     string
	CreatedAt string
	UpdatedAt string
}

func toDomain(rec Trainer) domain.Trainer {
	return domain.Trainer{
		ID:        rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		DOB:       rec.DOB,
		Gender:    rec.Gender,
		Phone:     rec.Phone,
		Address:   rec.Address,
		Picture:   rec.Picture,
		Field:     rec.Field,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Trainer) Trainer {
	return Trainer{
		ID:        rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		DOB:       rec.DOB,
		Gender:    rec.Gender,
		Phone:     rec.Phone,
		Address:   rec.Address,
		Picture:   rec.Picture,
		Field:     rec.Field,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
