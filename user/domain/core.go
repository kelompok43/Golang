package domain

import "mime/multipart"

type User struct {
	ID          int
	Picture     multipart.File
	PictureLink string
	Name        string
	DOB         string
	Email       string
	Password    string
	Phone       string
	Address     string
	// Gender      string
	Status    string
	CreatedAt string
	UpdatedAt string
}
