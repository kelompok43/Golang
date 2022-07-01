package domain

type Service interface {
	InsertData(domain Transaction) (transactionObj Transaction, err error)
	GetAllData() (transactionObj []Transaction, err error)
	GetByID(id int) (transactionObj Transaction, err error)
	UpdateStatus(id int, domain Transaction) (transactionObj Transaction, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Create(domain Transaction) (transactionObj Transaction, err error)
	Update(id int, domain Transaction) (transactionObj Transaction, err error)
	Get() (transactionObj []Transaction, err error)
	GetByID(id int) (domain Transaction, err error)
	Delete(id int) (err error)
}
