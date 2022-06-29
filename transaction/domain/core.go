package domain

import (
	"mime/multipart"
)

type Transaction struct {
	ID              int
	UserID          int
	PaymentMethodID int
	TotalPrice      int
	Status          string
	Payment_Receipt multipart.File
	PictureLink     string
	CreatedAt       string
	UpdatedAt       string
}
