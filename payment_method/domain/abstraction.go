package domain

type Service interface {
	InsertData(domain PaymentMethod) (pmObj PaymentMethod, err error)
	GetAllData() (pmObj []PaymentMethod, err error)
	GetByID(id int) (pmObj PaymentMethod, err error)
	UpdateData(id int, domain PaymentMethod) (pmObj PaymentMethod, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Create(domain PaymentMethod) (pmObj PaymentMethod, err error)
	Update(id int, domain PaymentMethod) (pmObj PaymentMethod, err error)
	Get() (pmObj []PaymentMethod, err error)
	GetByID(id int) (domain PaymentMethod, err error)
	Delete(id int) (err error)
}
