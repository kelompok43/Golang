package domain

type Membership struct {
	ID        int
	Category  string
	Price     int
	Duration  int
	CreatedAt string
	UpdatedAt string
}

type MembershipOrder struct {
	ID            int
	TransactionID int
	MembershipID  int
	Expired       string
	CreatedAt     string
	UpdatedAt     string
}
