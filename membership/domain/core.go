package domain

type MembershipCategory struct {
	ID        int
	Category  string
	Price     int
	Duration  int
	CreatedAt string
	UpdatedAt string
}

type Membership struct {
	ID                   int
	UserID               int
	MembershipCategoryID int
	ExpiredAt            string
	CreatedAt            string
	UpdatedAt            string
}
