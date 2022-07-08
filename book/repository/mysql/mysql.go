package repoMySQL

import (
	"github.com/kelompok43/Golang/book/domain"
	"gorm.io/gorm"
)

type bookRepository struct {
	DB *gorm.DB
}

// CreateOffline implements domain.Repository
func (br bookRepository) CreateOfflineClass(domain domain.OfflineClass) (categoryObj domain.OfflineClass, err error) {
	newRecord := fromOfflineClassDomain(domain)
	err = br.DB.Create(&newRecord).Error

	if err != nil {
		return categoryObj, err
	}

	categoryObj = toOfflineClassDomain(newRecord)
	return categoryObj, nil
}

// DeleteOffline implements domain.Repository
func (br bookRepository) DeleteOfflineClass(id int) (err error) {
	var record BookOfflineClass
	return br.DB.Delete(&record, id).Error
}

// GetCaOfflineID implements domain.Repository
func (br bookRepository) GetOfflineClassByID(id int) (offlineClassObj domain.OfflineClass, err error) {
	var newRecord BookOfflineClass
	err = br.DB.First(&newRecord, id).Error

	if err != nil {
		return offlineClassObj, err
	}

	return toOfflineClassDomain(newRecord), nil
}

// GetOffline implements domain.Repository
func (br bookRepository) GetOfflineClass() (offlineClassObj []domain.OfflineClass, err error) {
	var newRecords []BookOfflineClass

	err = br.DB.Find(&newRecords).Error

	if err != nil {
		return offlineClassObj, err
	}

	for _, value := range newRecords {
		offlineClassObj = append(offlineClassObj, toOfflineClassDomain(value))
	}

	return offlineClassObj, nil
}

// UpdateOffline implements domain.Repository
func (br bookRepository) UpdateOfflineClass(id int, domain domain.OfflineClass) (offlineClassObj domain.OfflineClass, err error) {
	var newRecord BookOfflineClass
	record := fromOfflineClassDomain(domain)
	err = br.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":               id,
		"user_id":          domain.UserID,
		"offline_class_id": domain.OfflineClassID,
		"created_at":       record.CreatedAt,
		"updated_at":       record.UpdatedAt,
	}).Error

	if err != nil {
		return offlineClassObj, err
	}

	offlineClassObj = toOfflineClassDomain(newRecord)
	return offlineClassObj, nil
}

// CreateOnlineClass implements domain.Repository
func (br bookRepository) CreateOnlineClass(domain domain.OnlineClass) (categoryObj domain.OnlineClass, err error) {
	newRecord := fromOnlineClassDomain(domain)
	err = br.DB.Create(&newRecord).Error

	if err != nil {
		return categoryObj, err
	}

	categoryObj = toOnlineClassDomain(newRecord)
	return categoryObj, nil
}

// DeleteOnlineClass implements domain.Repository
func (br bookRepository) DeleteOnlineClass(id int) (err error) {
	var record BookOnlineClass
	return br.DB.Delete(&record, id).Error
}

// GetCaOnlineClassID implements domain.Repository
func (br bookRepository) GetOnlineClassByID(id int) (onlineClassObj domain.OnlineClass, err error) {
	var newRecord BookOnlineClass
	err = br.DB.First(&newRecord, id).Error

	if err != nil {
		return onlineClassObj, err
	}

	return toOnlineClassDomain(newRecord), nil
}

// GetOnlineClass implements domain.Repository
func (br bookRepository) GetOnlineClass() (onlineClassObj []domain.OnlineClass, err error) {
	var newRecords []BookOnlineClass

	err = br.DB.Find(&newRecords).Error

	if err != nil {
		return onlineClassObj, err
	}

	for _, value := range newRecords {
		onlineClassObj = append(onlineClassObj, toOnlineClassDomain(value))
	}

	return onlineClassObj, nil
}

// UpdateOnlineClass implements domain.Repository
func (br bookRepository) UpdateOnlineClass(id int, domain domain.OnlineClass) (onlineClassObj domain.OnlineClass, err error) {
	var newRecord BookOnlineClass
	record := fromOnlineClassDomain(domain)
	err = br.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":              id,
		"user_id":         record.UserID,
		"online_class_id": record.OnlineClassID,
		"created_at":      record.CreatedAt,
		"updated_at":      record.UpdatedAt,
	}).Error

	if err != nil {
		return onlineClassObj, err
	}

	onlineClassObj = toOnlineClassDomain(newRecord)
	return onlineClassObj, nil
}

func NewBookRepository(db *gorm.DB) domain.Repository {
	return bookRepository{
		DB: db,
	}
}
