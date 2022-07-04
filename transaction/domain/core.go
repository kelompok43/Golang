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

type TransactionDetail struct {
	ID                   int
	TransactionID        int
	MembershipCategoryID int
	CreatedAt            string
	UpdatedAt            string
}
