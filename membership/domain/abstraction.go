package domain

type Service interface {
	InsertData(userID, categoryID int) (membershipObj Membership, err error)
	InsertCategory(domain MembershipCategory) (membershipCategoryObj MembershipCategory, err error)
	GetAllData() (membershipObj []Membership, err error)
	GetByID(id int) (membershipObj Membership, err error)
	GetAllCategory() (membershipCategoryObj []MembershipCategory, err error)
	GetCategoryByID(id int) (membershipCategoryObj MembershipCategory, err error)
	GetByUserID(userID int) (membershipObj Membership, err error)
	UpdateCategory(id int, domain MembershipCategory) (membershipCategoryObj MembershipCategory, err error)
	DeleteCategory(id int) (err error)
	GetCategoryByPrice(price int) (membershipCategoryObj MembershipCategory, err error)
}

type Repository interface {
	Create(domain Membership) (membershipObj Membership, err error)
	CreateCategory(domain MembershipCategory) (membershipCategoryObj MembershipCategory, err error)
	UpdateCategory(id int, domain MembershipCategory) (membershipCategoryObj MembershipCategory, err error)
	Get() (membershipObj []Membership, err error)
	GetByID(id int) (membershipObj Membership, err error)
	Update(id int, domain Membership) (membershipObj Membership, err error)
	GetCategory() (membershipCategoryObj []MembershipCategory, err error)
	GetCategoryByID(id int) (membershipCategoryObj MembershipCategory, err error)
	GetCategoryByPrice(price int) (membershipCategoryObj MembershipCategory, err error)
	GetByUserID(userID int) (membershipObj Membership, err error)
	// GetByUserID(userID int) (membershipOrderObj []MembershipOrder, err error) //ini di service transaction
	DeleteCategory(id int) (err error)
}
