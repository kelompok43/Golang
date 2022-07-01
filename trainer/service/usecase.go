package service

import (
	"bytes"
	"errors"
	"io"

	storageHelper "github.com/kelompok43/Golang/helpers/azure"
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

	emailCheck, _ := ts.repository.GetByEmail(domain.Email)
	if trainer.Email != emailCheck.Email {
		_, errGetTrainer := ts.repository.GetByEmail(domain.Email)
		if errGetTrainer == nil {
			return trainerObj, errors.New("email telah terdaftar")
		}
	}

	domain.CreatedAt = trainer.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	domain.PictureLink = trainer.PictureLink
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

	// sPicture := fmt.Sprintf("%v", domain.Picture)
	// data, _ := ioutil.ReadFile(sPicture)
	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, domain.Picture); err != nil {
		return trainerObj, err
	}

	data := buf.Bytes()
	domain.PictureLink, _ = storageHelper.UploadBytesToBlob(data)
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
