package repoMySQL

import (
	"github.com/kelompok43/Golang/membership/domain"
	repoMYSQLTrx "github.com/kelompok43/Golang/transaction/repository/mysql"
	"gorm.io/gorm"
)

type MembershipCategory struct {
	gorm.Model
	ID          int
	Category    string
	Price       int
	Duration    int
	CreatedAt   string
	UpdatedAt   string
	Membership  Membership
	Transaction repoMYSQLTrx.TransactionDetail `gorm:"foreignKey:MembershipCategoryID"`
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

type Membership struct {
	gorm.Model
	ID                   int
	UserID               int
	MembershipCategoryID int
	ExpiredAt            string
	CreatedAt            string
	UpdatedAt            string
}

func toDomain(rec Membership) domain.Membership {
	return domain.Membership{
		ID:                   rec.ID,
		UserID:               rec.UserID,
		MembershipCategoryID: rec.MembershipCategoryID,
		ExpiredAt:            rec.ExpiredAt,
		CreatedAt:            rec.CreatedAt,
		UpdatedAt:            rec.UpdatedAt,
	}
}

func categoryToDomain(rec MembershipCategory) domain.MembershipCategory {
	return domain.MembershipCategory{
		ID:        rec.ID,
		Category:  rec.Category,
		Price:     rec.Price,
		Duration:  rec.Duration,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Membership) Membership {
	return Membership{
		ID:                   rec.ID,
		UserID:               rec.UserID,
		MembershipCategoryID: rec.MembershipCategoryID,
		ExpiredAt:            rec.ExpiredAt,
		CreatedAt:            rec.CreatedAt,
		UpdatedAt:            rec.UpdatedAt,
	}
}

func fromDomainToCategory(rec domain.MembershipCategory) MembershipCategory {
	return MembershipCategory{
		ID:        rec.ID,
		Category:  rec.Category,
		Price:     rec.Price,
		Duration:  rec.Duration,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
