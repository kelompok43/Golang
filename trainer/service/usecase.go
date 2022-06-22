package service

import (
	"errors"
	"fmt"

	timeHelper "github.com/kelompok43/Golang/helpers/time"
	"github.com/kelompok43/Golang/trainer/domain"
)

type trainerService struct {
	repository domain.Repository
}

// DeleteData implements domain.Service
func (ts trainerService) DeleteData(id int) (err error) {
	errResp := ts.repository.Delete(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// UpdateData implements domain.Service
func (ts trainerService) UpdateData(id int, domain domain.Trainer) (trainerObj domain.Trainer, err error) {
	trainer, errGetByID := ts.GetByID(id)

	if errGetByID != nil {
		return trainer, errGetByID
	}

	fmt.Println(trainer)

	domain.CreatedAt = trainer.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	trainerObj, err = ts.repository.Update(id, domain)

	if err != nil {
		return trainerObj, err
	}

	return trainerObj, nil
}

// GetByEmail implements domain.Service
func (ts trainerService) GetByEmail(email string) (trainerObj domain.Trainer, err error) {
	trainerObj, err = ts.repository.GetByEmail(email)

	if err != nil {
		return trainerObj, err
	}

	return trainerObj, nil
}

// GetByID implements domain.Service
func (ts trainerService) GetByID(id int) (trainerObj domain.Trainer, err error) {
	trainerObj, err = ts.repository.GetByID(id)

	if err != nil {
		return trainerObj, err
	}

	return trainerObj, nil
}

func (ts trainerService) InsertData(domain domain.Trainer) (trainerObj domain.Trainer, err error) {
	email := domain.Email
	_, errGetTrainer := ts.repository.GetByEmail(email)

	if errGetTrainer == nil {
		return trainerObj, errors.New("email telah terdaftar")
	}

	if err != nil {
		return trainerObj, err
	}

	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	trainerObj, err = ts.repository.Create(domain)

	if err != nil {
		return trainerObj, err
	}

	return trainerObj, nil
}

// GetAllData implements domain.Service
func (ts trainerService) GetAllData() (trainerObj []domain.Trainer, err error) {
	trainerObj, _ = ts.repository.Get()

	if err != nil {
		return trainerObj, err
	}

	return trainerObj, nil
}

func NewTrainerService(repo domain.Repository) domain.Service {
	return trainerService{
		repository: repo,
	}
}
