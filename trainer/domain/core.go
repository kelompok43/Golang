package domain

import (
	"mime/multipart"
)

type Trainer struct {
	ID          int
	Name        string
	Email       string
	DOB         string
	Gender      string
	Phone       string
	Address     string
	Picture     multipart.File
	Field       string
	PictureLink string
	CreatedAt   string
	UpdatedAt   string
}
