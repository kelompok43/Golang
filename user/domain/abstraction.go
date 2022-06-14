package domain

type Service interface {
	CreateToken(email, password string) (token string, userObj User, err error)
	InsertData(domain User) (userObj User, err error)
	InsertDetailData(domain User) (userObj User, err error)
	GetAllData() (userObj []User, err error)
	GetByID(id int) (userObj User, err error)
	GetByEmail(email string) (userObj User, err error)
	ChangePassword(id int, domain User) (userObj User, err error)
}

type Repository interface {
	Create(domain User) (userObj User, err error)
	Update(domain User) (userObj User, err error)
	AddDetail(domain User) (userObj User, err error)
	Get() (userObj []User, err error)
	GetByID(id int) (domain User, err error)
	GetByEmail(email string) (userObj User, err error)
	GetDetail(id int) (userObj User, err error)
}
