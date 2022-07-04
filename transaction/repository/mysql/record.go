package repoMySQL

import (
	"github.com/kelompok43/Golang/transaction/domain"
	"gorm.io/gorm"
)

// type Transaction struct {
// 	gorm.Model
// 	ID                int
// 	UserID            int
// 	PaymentMethodID   int
// 	TotalPrice        int
// 	Status            string
// 	PictureLink       string
// 	CreatedAt         string
// 	UpdatedAt         string
// 	Memberships       []repoMySQLM.MembershipOrder `gorm:"foreignKey:TransactionID"`
// 	TransactionDetail TransactionDetail
// }

type Transaction struct {
	gorm.Model
	ID                int
	UserID            int
	PaymentMethodID   int
	TotalPrice        int
	Status            string
	PictureLink       string
	CreatedAt         string
	UpdatedAt         string
	TransactionDetail TransactionDetail
}

type TransactionDetail struct {
	gorm.Model
	ID                   int
	TransactionID        int
	MembershipCategoryID int
	CreatedAt            string
	UpdatedAt            string
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

func detailToDomain(rec TransactionDetail) domain.TransactionDetail {
	return domain.TransactionDetail{
		ID:                   rec.ID,
		TransactionID:        rec.TransactionID,
		MembershipCategoryID: rec.MembershipCategoryID,
		CreatedAt:            rec.CreatedAt,
		UpdatedAt:            rec.UpdatedAt,
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

func fromDomainToDetail(rec domain.TransactionDetail) TransactionDetail {
	return TransactionDetail{
		ID:                   rec.ID,
		TransactionID:        rec.TransactionID,
		MembershipCategoryID: rec.MembershipCategoryID,
		CreatedAt:            rec.CreatedAt,
		UpdatedAt:            rec.UpdatedAt,
	}
}
