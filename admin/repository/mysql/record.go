package repoMySQL

import (
	"github.com/kelompok43/Golang/admin/domain"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	ID        int
	Name      string
	DOB       string
	Gender    string
	Address   string
	Role      string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func toDomain(rec Admin) domain.Admin {
	return domain.Admin{
		ID:        rec.ID,
		Name:      rec.Name,
		DOB:       rec.DOB,
		Address:   rec.Address,
		Gender:    rec.Gender,
		Role:      rec.Role,
		Email:     rec.Email,
		Password:  rec.Password,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Admin) Admin {
	return Admin{
		ID:        rec.ID,
		Name:      rec.Name,
		DOB:       rec.DOB,
		Address:   rec.Address,
		Gender:    rec.Gender,
		Role:      rec.Role,
		Email:     rec.Email,
		Password:  rec.Password,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
