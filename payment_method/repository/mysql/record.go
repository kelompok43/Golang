package repoMySQL

import (
	"github.com/kelompok43/Golang/payment_method/domain"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	ID        int
	Name      string
	AccNumber string
	AccName   string
	CreatedAt string
	UpdatedAt string
}

func toDomain(rec PaymentMethod) domain.PaymentMethod {
	return domain.PaymentMethod{
		ID:        rec.ID,
		Name:      rec.Name,
		AccNumber: rec.AccNumber,
		AccName:   rec.AccName,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(rec domain.PaymentMethod) PaymentMethod {
	return PaymentMethod{
		ID:        rec.ID,
		Name:      rec.Name,
		AccNumber: rec.AccNumber,
		AccName:   rec.AccName,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
