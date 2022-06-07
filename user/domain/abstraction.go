package domain

type Service interface {
	CreateToken(email, password string) (token string, err error)
	InsertData(domain User) (userObj User, err error)
	GetByEmailPassword(email, password string) (id int, status string, err error)
}

type Repository interface {
	Create(domain User) (userObj User, err error)
	GetByEmailPassword(email, password string) (domain User, err error)
}
