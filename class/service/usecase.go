package service

import (
	"bytes"
	"io"

	"github.com/kelompok43/Golang/class/domain"
	storageHelper "github.com/kelompok43/Golang/helpers/azure"
	timeHelper "github.com/kelompok43/Golang/helpers/time"
)

type classService struct {
	repository domain.Repository
}

// DeleteOffline implements domain.Service
func (cs classService) DeleteOffline(id int) (err error) {
	errResp := cs.repository.DeleteOffline(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// GetAllOffline implements domain.Service
func (cs classService) GetAllOffline() (offlineClassObj []domain.Offline, err error) {
	offlineClassObj, _ = cs.repository.GetOffline()

	if err != nil {
		return offlineClassObj, err
	}

	return offlineClassObj, nil
}

// GetOfflineByID implements domain.Service
func (cs classService) GetOfflineByID(id int) (offlineClassObj domain.Offline, err error) {
	offlineClassObj, err = cs.repository.GetOfflineByID(id)

	if err != nil {
		return offlineClassObj, err
	}

	return offlineClassObj, nil
}

// InsertOffline implements domain.Service
func (cs classService) InsertOffline(domain domain.Offline) (offlineClassObj domain.Offline, err error) {
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	offlineClassObj, err = cs.repository.CreateOffline(domain)

	if err != nil {
		return offlineClassObj, err
	}

	return offlineClassObj, nil
}

// UpdateOffline implements domain.Service
func (cs classService) UpdateOffline(id int, domain domain.Offline) (offlineClassObj domain.Offline, err error) {
	offlineClass, errGetByID := cs.GetOfflineByID(id)

	if errGetByID != nil {
		return offlineClass, errGetByID
	}

	domain.CreatedAt = offlineClass.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	offlineClassObj, err = cs.repository.UpdateOffline(id, domain)

	if err != nil {
		return offlineClassObj, err
	}

	return offlineClassObj, nil
}

// DeleteOnline implements domain.Service
func (cs classService) DeleteOnline(id int) (err error) {
	errResp := cs.repository.DeleteOnline(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// GetAllOnline implements domain.Service
func (cs classService) GetAllOnline() (onlineClassObj []domain.Online, err error) {
	onlineClassObj, _ = cs.repository.GetOnline()

	if err != nil {
		return onlineClassObj, err
	}

	return onlineClassObj, nil
}

// GetOnlineByID implements domain.Service
func (cs classService) GetOnlineByID(id int) (onlineClassObj domain.Online, err error) {
	onlineClassObj, err = cs.repository.GetOnlineByID(id)

	if err != nil {
		return onlineClassObj, err
	}

	return onlineClassObj, nil
}

// InsertOnline implements domain.Service
func (cs classService) InsertOnline(domain domain.Online) (onlineClassObj domain.Online, err error) {
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	onlineClassObj, err = cs.repository.CreateOnline(domain)

	if err != nil {
		return onlineClassObj, err
	}

	return onlineClassObj, nil
}

// UpdateOnline implements domain.Service
func (cs classService) UpdateOnline(id int, domain domain.Online) (onlineClassObj domain.Online, err error) {
	onlineClass, errGetByID := cs.GetOnlineByID(id)

	if errGetByID != nil {
		return onlineClass, errGetByID
	}

	domain.CreatedAt = onlineClass.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	onlineClassObj, err = cs.repository.UpdateOnline(id, domain)

	if err != nil {
		return onlineClassObj, err
	}

	return onlineClassObj, nil
}

// DeleteCategory implements domain.Service
func (cs classService) DeleteCategory(id int) (err error) {
	errResp := cs.repository.DeleteCategory(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// UpdateCategory implements domain.Service
func (cs classService) UpdateCategory(id int, domain domain.Category) (categoryObj domain.Category, err error) {
	category, errGetByID := cs.GetCategoryByID(id)

	if errGetByID != nil {
		return category, errGetByID
	}

	domain.CreatedAt = category.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	domain.PictureLink = category.PictureLink
	categoryObj, err = cs.repository.UpdateCategory(id, domain)

	if err != nil {
		return categoryObj, err
	}

	return categoryObj, nil
}

// GetCategoryByID implements domain.Service
func (cs classService) GetCategoryByID(id int) (categoryObj domain.Category, err error) {
	categoryObj, err = cs.repository.GetCategoryByID(id)

	if err != nil {
		return categoryObj, err
	}

	return categoryObj, nil
}

func (cs classService) InsertCategory(domain domain.Category) (categoryObj domain.Category, err error) {
	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, domain.Picture); err != nil {
		return categoryObj, err
	}

	data := buf.Bytes()
	domain.PictureLink, _ = storageHelper.UploadBytesToBlob(data)
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	categoryObj, err = cs.repository.CreateCategory(domain)

	if err != nil {
		return categoryObj, err
	}

	return categoryObj, nil
}

// GetAllCategory implements domain.Service
func (cs classService) GetAllCategory() (categoryObj []domain.Category, err error) {
	categoryObj, _ = cs.repository.GetCategory()

	if err != nil {
		return categoryObj, err
	}

	return categoryObj, nil
}

func NewClassService(repo domain.Repository) domain.Service {
	return classService{
		repository: repo,
	}
}
