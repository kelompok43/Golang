package domain

type Service interface {
	CreateToken(email, password string) (token string, adminObj Admin, err error)
	InsertData(domain Admin) (adminObj Admin, err error)
	GetAllData() (adminObj []Admin, err error)
	GetByID(id int) (adminObj Admin, err error)
	GetByEmail(email string) (adminObj Admin, err error)
	ChangePassword(id int, domain Admin) (adminObj Admin, err error)
	UpdateData(id int, domain Admin) (adminObj Admin, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Create(domain Admin) (adminObj Admin, err error)
	Update(domain Admin) (adminObj Admin, err error)
	Get() (adminObj []Admin, err error)
	GetByID(id int) (domain Admin, err error)
	GetByEmail(email string) (adminObj Admin, err error)
	Delete(id int) (err error)
}
