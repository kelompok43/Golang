package repoMySQL

import (
	repoMySQLC "github.com/kelompok43/Golang/class/repository/mysql"
	"github.com/kelompok43/Golang/trainer/domain"
	"gorm.io/gorm"
)

type Trainer struct {
	gorm.Model
	ID             int
	Name           string
	Email          string
	DOB            string
	Gender         string
	Phone          string
	Address        string
	PictureLink    string
	Field          string
	CreatedAt      string
	UpdatedAt      string
	OnlineClasses  []repoMySQLC.OnlineClass  `gorm:"foreignKey:TrainerID"`
	OfflineClasses []repoMySQLC.OfflineClass `gorm:"foreignKey:TrainerID"`
}

func toDomain(rec Trainer) domain.Trainer {
	return domain.Trainer{
		ID:          rec.ID,
		Name:        rec.Name,
		Email:       rec.Email,
		DOB:         rec.DOB,
		Gender:      rec.Gender,
		Phone:       rec.Phone,
		Address:     rec.Address,
		PictureLink: rec.PictureLink,
		Field:       rec.Field,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Trainer) Trainer {
	return Trainer{
		ID:          rec.ID,
		Name:        rec.Name,
		Email:       rec.Email,
		DOB:         rec.DOB,
		Gender:      rec.Gender,
		Phone:       rec.Phone,
		Address:     rec.Address,
		PictureLink: rec.PictureLink,
		Field:       rec.Field,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}
