package repoMySQL

import (
	"github.com/kelompok43/Golang/user/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

// Create implements domain.Repository
func (ur userRepository) Create(domain domain.User) (userObj domain.User, err error) {
	panic("unimplemented")
}

// GetByEmailPassword implements domain.Repository
func (ur userRepository) GetByEmailPassword(email string, password string) (domain domain.User, err error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return userRepository{
		DB: db,
	}
}
