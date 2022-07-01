package repoMySQL

import (
	"github.com/kelompok43/Golang/membership/domain"
	"gorm.io/gorm"
)

type Membership struct {
	gorm.Model
	ID               int
	Category         string
	Price            int
	Duration         int
	CreatedAt        string
	UpdatedAt        string
	MembershipOrders []MembershipOrder
}

type MembershipOrder struct {
	gorm.Model
	ID            int
	TransactionID int
	MembershipID  int
	Expired       string
	CreatedAt     string
	UpdatedAt     string
}

func toDomain(rec Membership) domain.Membership {
	return domain.Membership{
		ID:        rec.ID,
		Category:  rec.Category,
		Price:     rec.Price,
		Duration:  rec.Duration,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func orderToDomain(rec MembershipOrder) domain.MembershipOrder {
	return domain.MembershipOrder{
		ID:            rec.ID,
		TransactionID: rec.TransactionID,
		MembershipID:  rec.MembershipID,
		Expired:       rec.Expired,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Membership) Membership {
	return Membership{
		ID:        rec.ID,
		Category:  rec.Category,
		Price:     rec.Price,
		Duration:  rec.Duration,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomainToOrder(rec domain.MembershipOrder) MembershipOrder {
	return MembershipOrder{
		ID:            rec.ID,
		TransactionID: rec.TransactionID,
		MembershipID:  rec.MembershipID,
		Expired:       rec.Expired,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
}
