package repoMySQL

import "gorm.io/gorm"

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
	CreatedAt string
	UpdatedAt string
}
