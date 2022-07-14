package domain

import (
	"mime/multipart"
)

type Category struct {
	ID          int
	Name        string
	Description string
	Picture     multipart.File
	PictureLink string
	CreatedAt   string
	UpdatedAt   string
}

type Online struct {
	ID              int
	ClassCategoryID int
	TrainerID       int
	Date            string
	StartedAt       string
	EndedAt         string
	Link            string
	CreatedAt       string
	UpdatedAt       string
}

type Offline struct {
	ID              int
	ClassCategoryID int
	TrainerID       int
	Date            string
	StartedAt       string
	EndedAt         string
	Place           string
	Quota           int
	CreatedAt       string
	UpdatedAt       string
}
