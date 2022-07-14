package domain

type Service interface {
	InsertData(domain Transaction) (transactionObj Transaction, err error)
	GetAllData() (transactionObj []Transaction, err error)
	GetByID(id int) (transactionObj Transaction, err error)
	GetUserTrx(userID int) (transactionObj []Transaction, err error)
	GetUserTrxByID(userID, trxID int) (transactionObj Transaction, err error)
	UpdateStatus(id int, domain Transaction) (transactionObj Transaction, err error)
	DeleteData(id int) (err error)
	InsertDetail(id, price int) (transactionDetailObj TransactionDetail, err error)
}

type Repository interface {
	Create(domain Transaction) (transactionObj Transaction, err error)
	CreateDetail(domain TransactionDetail) (transactionDetailObj TransactionDetail, err error)
	Update(id int, domain Transaction) (transactionObj Transaction, err error)
	Get() (transactionObj []Transaction, err error)
	GetByID(id int) (transactionObj Transaction, err error)
	GetByUserID(userID int) (transactionObj []Transaction, err error)
	Delete(id int) (err error)
}
