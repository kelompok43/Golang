package domain

type Service interface {
	InsertData(domain Trainer) (trainerObj Trainer, err error)
	GetAllData() (trainerObj []Trainer, err error)
	GetByID(id int) (trainerObj Trainer, err error)
	GetByEmail(email string) (trainerObj Trainer, err error)
	UpdateData(id int, domain Trainer) (trainerObj Trainer, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Create(domain Trainer) (trainerObj Trainer, err error)
	Update(id int, domain Trainer) (trainerObj Trainer, err error)
	Get() (trainerObj []Trainer, err error)
	GetByID(id int) (domain Trainer, err error)
	GetByEmail(email string) (trainerObj Trainer, err error)
	Delete(id int) (err error)
}
