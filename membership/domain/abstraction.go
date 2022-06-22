package domain

type Service interface {
	InsertData(domain Membership) (pmObj Membership, err error)
	GetAllData() (pmObj []Membership, err error)
	GetByID(id int) (pmObj Membership, err error)
	UpdateData(id int, domain Membership) (pmObj Membership, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Create(domain Membership) (pmObj Membership, err error)
	Update(id int, domain Membership) (pmObj Membership, err error)
	Get() (pmObj []Membership, err error)
	GetByID(id int) (domain Membership, err error)
	Delete(id int) (err error)
}
