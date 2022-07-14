package repoMySQL

import (
	repoMYSQLB "github.com/kelompok43/Golang/book/repository/mysql"
	repoMYSQLM "github.com/kelompok43/Golang/membership/repository/mysql"
	repoMYSQLTrx "github.com/kelompok43/Golang/transaction/repository/mysql"
	"github.com/kelompok43/Golang/user/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                 int
	Name               string
	Email              string
	Password           string
	Status             string
	CreatedAt          string
	UpdatedAt          string
	UserDetail         UserDetail
	Transactions       []repoMYSQLTrx.Transaction    `gorm:"foreignKey:UserID"`
	Membership         repoMYSQLM.Membership         `gorm:"foreignKey:UserID"`
	BookOfflineClasses []repoMYSQLB.BookOfflineClass `gorm:"foreignKey:UserID"`
	BookOnlineClasses  []repoMYSQLB.BookOnlineClass  `gorm:"foreignKey:UserID"`
}

type UserDetail struct {
	gorm.Model
	UserID  int
	DOB     string
	Phone   string
	Address string
	// Gender      string
	PictureLink string
	CreatedAt   string
	UpdatedAt   string
	// CreatedAt string `gorm:"autoCreateTime:false"`
	// UpdatedAt string `gorm:"autoCreateTime:false"`
}

type joinResult struct {
	ID       int
	Name     string
	DOB      string
	Email    string
	Password string
	Phone    string
	Address  string
	// Gender    string
	PictureLink string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

func toDomain(rec joinResult) domain.User {
	return domain.User{
		ID:       rec.ID,
		Name:     rec.Name,
		DOB:      rec.DOB,
		Email:    rec.Email,
		Password: rec.Password,
		Phone:    rec.Phone,
		Address:  rec.Address,
		// Gender:    rec.Gender,
		PictureLink: rec.PictureLink,
		Status:      rec.Status,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
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
		UserID:      rec.ID,
		DOB:         rec.DOB,
		Phone:       rec.Phone,
		Address:     rec.Address,
		PictureLink: rec.PictureLink,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}
