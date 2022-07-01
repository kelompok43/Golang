package repoMySQL

import (
	repoMySQLM "github.com/kelompok43/Golang/membership/repository/mysql"
	"github.com/kelompok43/Golang/transaction/domain"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID              int
	UserID          int
	PaymentMethodID int
	TotalPrice      int
	Status          string
	PictureLink     string
	CreatedAt       string
	UpdatedAt       string
	Memberships     []repoMySQLM.MembershipOrder `gorm:"foreignKey:TransactionID"`
}

func toDomain(rec Transaction) domain.Transaction {
	return domain.Transaction{
		ID:              rec.ID,
		UserID:          rec.UserID,
		PaymentMethodID: rec.PaymentMethodID,
		TotalPrice:      rec.TotalPrice,
		Status:          rec.Status,
		PictureLink:     rec.PictureLink,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Transaction) Transaction {
	return Transaction{
		ID:              rec.ID,
		UserID:          rec.UserID,
		PaymentMethodID: rec.PaymentMethodID,
		TotalPrice:      rec.TotalPrice,
		Status:          rec.Status,
		PictureLink:     rec.PictureLink,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
	}
}
