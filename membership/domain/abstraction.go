package domain

type Service interface {
	InsertData(domain Membership) (membershipObj Membership, err error)
	GetAllData() (membershipObj []Membership, err error)
	GetByID(id int) (membershipObj Membership, err error)
	UpdateData(id int, domain Membership) (membershipObj Membership, err error)
	DeleteData(id int) (err error)
	InsertOrder(transactionID, price int) (membershipOrderObj MembershipOrder, err error)
	GetOrderByUserID(userID int) (membershipOrderObj []MembershipOrder, err error)
}

type Repository interface {
	Create(domain Membership) (membershipObj Membership, err error)
	CreateOrder(domain MembershipOrder) (membershipOrderObj MembershipOrder, err error)
	Update(id int, domain Membership) (membershipObj Membership, err error)
	Get() (membershipObj []Membership, err error)
	GetByID(id int) (membershipObj Membership, err error)
	GetByPrice(price int) (membershipObj Membership, err error)
	GetByUserID(userID int) (membershipOrderObj []MembershipOrder, err error)
	Delete(id int) (err error)
}
