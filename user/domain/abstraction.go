package domain

type Service interface {
	CreateToken(email, password string) (token string, err error)
	InsertData(domain User) (userObj User, err error)
	GetAllData() (userObj []User, err error)
	GetByID(id int) (userObj User, err error)
	GetByEmailPassword(email, password string) (id int, status string, err error)
}

type Repository interface {
	Create(domain User) (userObj User, err error)
	Get() (userObj []User, err error)
	GetByEmail(email string) (userObj User, err error)
	GetByEmailPassword(email, password string) (domain User, err error)
	GetByID(id int) (domain User, err error)
}
