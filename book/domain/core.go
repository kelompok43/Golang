package domain

type OnlineClass struct {
	ID            int
	UserID        int
	OnlineClassID int
	CreatedAt     string
	UpdatedAt     string
}

type OfflineClass struct {
	ID             int
	UserID         int
	OfflineClassID int
	CreatedAt      string
	UpdatedAt      string
}
