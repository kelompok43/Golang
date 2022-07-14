package domain

import (
	"mime/multipart"
)

type News struct {
	ID             int
	NewsCategoryID int
	Title          string
	Author         string
	Date           string
	Description    string
	Picture        multipart.File
	PictureLink    string
	CreatedAt      string
	UpdatedAt      string
}

type Category struct {
	ID        int
	Name      string
	CreatedAt string
	UpdatedAt string
}
