package repoMySQL

import (
	"github.com/kelompok43/Golang/user/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         int
	Name       string
	Email      string
	Password   string
	Status     string
	CreatedAt  string
	UpdatedAt  string
	UserDetail UserDetail
}

type UserDetail struct {
	gorm.Model
	UserID    int
	DOB       string
	Phone     string
	Address   string
	Gender    string
	CreatedAt string `gorm:"autoCreateTime:false"`
	UpdatedAt string `gorm:"autoCreateTime:false"`
}

type joinResult struct {
	ID        int
	Name      string
	DOB       string
	Email     string
	Password  string
	Phone     string
	Address   string
	Gender    string
	Status    string
	CreatedAt string
	UpdatedAt string
}

func toDomain(rec joinResult) domain.User {
	return domain.User{
		ID:        rec.ID,
		Name:      rec.Name,
		DOB:       rec.DOB,
		Email:     rec.Email,
		Password:  rec.Password,
		Phone:     rec.Phone,
		Address:   rec.Address,
		Gender:    rec.Gender,
		Status:    rec.Status,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomainToUser(rec domain.User) User {
	return User{
		ID:        rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Status:    rec.Status,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomainToUserDetail(rec domain.User) UserDetail {
	return UserDetail{
		UserID:    rec.ID,
		DOB:       rec.DOB,
		Phone:     rec.Phone,
		Address:   rec.Address,
		Gender:    rec.Gender,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
