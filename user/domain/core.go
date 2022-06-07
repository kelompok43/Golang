package domain

import "time"

type User struct {
	ID        int
	Name      string
	DOB       string
	Email     string
	Password  string
	Phone     string
	Address   string
	Gender    string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
