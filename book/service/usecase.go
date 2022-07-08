package service

import (
	"github.com/kelompok43/Golang/book/domain"
	timeHelper "github.com/kelompok43/Golang/helpers/time"
)

type bookService struct {
	repository domain.Repository
}

// DeleteOfflineClass implements domain.Service
func (bs bookService) DeleteOfflineClass(id int) (err error) {
	errResp := bs.repository.DeleteOfflineClass(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// GetAllOfflineClass implements domain.Service
func (bs bookService) GetAllOfflineClass() (offlineClassObj []domain.OfflineClass, err error) {
	offlineClassObj, _ = bs.repository.GetOfflineClass()

	if err != nil {
		return offlineClassObj, err
	}

	return offlineClassObj, nil
}

// GetOfflineClassByID implements domain.Service
func (bs bookService) GetOfflineClassByID(id int) (offlineClassObj domain.OfflineClass, err error) {
	offlineClassObj, err = bs.repository.GetOfflineClassByID(id)

	if err != nil {
		return offlineClassObj, err
	}

	return offlineClassObj, nil
}

// InsertOfflineClass implements domain.Service
func (bs bookService) InsertOfflineClass(domain domain.OfflineClass) (offlineClassObj domain.OfflineClass, err error) {
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	offlineClassObj, err = bs.repository.CreateOfflineClass(domain)

	if err != nil {
		return offlineClassObj, err
	}

	return offlineClassObj, nil
}

// UpdateOfflineClass implements domain.Service
func (bs bookService) UpdateOfflineClass(id int, domain domain.OfflineClass) (offlineClassObj domain.OfflineClass, err error) {
	offlineClass, errGetByID := bs.GetOfflineClassByID(id)

	if errGetByID != nil {
		return offlineClass, errGetByID
	}

	domain.CreatedAt = offlineClass.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	offlineClassObj, err = bs.repository.UpdateOfflineClass(id, domain)

	if err != nil {
		return offlineClassObj, err
	}

	return offlineClassObj, nil
}

// DeleteOnlineClass implements domain.Service
func (bs bookService) DeleteOnlineClass(id int) (err error) {
	errResp := bs.repository.DeleteOnlineClass(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// GetAllOnlineClass implements domain.Service
func (bs bookService) GetAllOnlineClass() (onlineClassObj []domain.OnlineClass, err error) {
	onlineClassObj, _ = bs.repository.GetOnlineClass()

	if err != nil {
		return onlineClassObj, err
	}

	return onlineClassObj, nil
}

// GetOnlineClassByID implements domain.Service
func (bs bookService) GetOnlineClassByID(id int) (onlineClassObj domain.OnlineClass, err error) {
	onlineClassObj, err = bs.repository.GetOnlineClassByID(id)

	if err != nil {
		return onlineClassObj, err
	}

	return onlineClassObj, nil
}

// InsertOnlineClass implements domain.Service
func (bs bookService) InsertOnlineClass(domain domain.OnlineClass) (onlineClassObj domain.OnlineClass, err error) {
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	onlineClassObj, err = bs.repository.CreateOnlineClass(domain)

	if err != nil {
		return onlineClassObj, err
	}

	return onlineClassObj, nil
}

// UpdateOnlineClass implements domain.Service
func (bs bookService) UpdateOnlineClass(id int, domain domain.OnlineClass) (onlineClassObj domain.OnlineClass, err error) {
	onlineClass, errGetByID := bs.GetOnlineClassByID(id)

	if errGetByID != nil {
		return onlineClass, errGetByID
	}

	domain.CreatedAt = onlineClass.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	onlineClassObj, err = bs.repository.UpdateOnlineClass(id, domain)

	if err != nil {
		return onlineClassObj, err
	}

	return onlineClassObj, nil
}

func NewBookService(repo domain.Repository) domain.Service {
	return bookService{
		repository: repo,
	}
}
