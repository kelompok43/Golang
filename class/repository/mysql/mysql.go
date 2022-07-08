package repoMySQL

import (
	"errors"

	"github.com/kelompok43/Golang/class/domain"
	"gorm.io/gorm"
)

type classRepository struct {
	DB *gorm.DB
}

// CreateOffline implements domain.Repository
func (cr classRepository) CreateOffline(domain domain.Offline) (categoryObj domain.Offline, err error) {
	newRecord := fromOfflineDomain(domain)
	err = cr.DB.Create(&newRecord).Error

	if err != nil {
		return categoryObj, err
	}

	categoryObj = toOfflineDomain(newRecord)
	return categoryObj, nil
}

// DeleteOffline implements domain.Repository
func (cr classRepository) DeleteOffline(id int) (err error) {
	var record OfflineClass
	return cr.DB.Delete(&record, id).Error
}

// GetCaOfflineID implements domain.Repository
func (cr classRepository) GetOfflineByID(id int) (offlineClassObj domain.Offline, err error) {
	var newRecord OfflineClass
	err = cr.DB.First(&newRecord, id).Error

	if err != nil {
		return offlineClassObj, err
	}

	return toOfflineDomain(newRecord), nil
}

// GetOffline implements domain.Repository
func (cr classRepository) GetOffline() (offlineClassObj []domain.Offline, err error) {
	var newRecords []OfflineClass

	err = cr.DB.Find(&newRecords).Error

	if err != nil {
		return offlineClassObj, err
	}

	for _, value := range newRecords {
		offlineClassObj = append(offlineClassObj, toOfflineDomain(value))
	}

	return offlineClassObj, nil
}

// UpdateOffline implements domain.Repository
func (cr classRepository) UpdateOffline(id int, domain domain.Offline) (offlineClassObj domain.Offline, err error) {
	var newRecord OfflineClass
	record := fromOfflineDomain(domain)
	err = cr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":                id,
		"trainer_id":        record.TrainerID,
		"class_category_id": record.ClassCategoryID,
		"date":              record.Date,
		"started_at":        record.StartedAt,
		"ended_at":          record.EndedAt,
		"place":             record.Place,
		"quota":             record.Quota,
		"created_at":        record.CreatedAt,
		"updated_at":        record.UpdatedAt,
	}).Error

	if err != nil {
		return offlineClassObj, err
	}

	offlineClassObj = toOfflineDomain(newRecord)
	return offlineClassObj, nil
}

// CreateOnline implements domain.Repository
func (cr classRepository) CreateOnline(domain domain.Online) (categoryObj domain.Online, err error) {
	newRecord := fromOnlineDomain(domain)
	err = cr.DB.Create(&newRecord).Error

	if err != nil {
		return categoryObj, err
	}

	categoryObj = toOnlineDomain(newRecord)
	return categoryObj, nil
}

// DeleteOnline implements domain.Repository
func (cr classRepository) DeleteOnline(id int) (err error) {
	var record OnlineClass
	return cr.DB.Delete(&record, id).Error
}

// GetCaOnlineID implements domain.Repository
func (cr classRepository) GetOnlineByID(id int) (onlineClassObj domain.Online, err error) {
	var newRecord OnlineClass
	err = cr.DB.First(&newRecord, id).Error

	if err != nil {
		return onlineClassObj, err
	}

	return toOnlineDomain(newRecord), nil
}

// GetOnline implements domain.Repository
func (cr classRepository) GetOnline() (onlineClassObj []domain.Online, err error) {
	var newRecords []OnlineClass

	err = cr.DB.Find(&newRecords).Error

	if err != nil {
		return onlineClassObj, err
	}

	for _, value := range newRecords {
		onlineClassObj = append(onlineClassObj, toOnlineDomain(value))
	}

	return onlineClassObj, nil
}

// UpdateOnline implements domain.Repository
func (cr classRepository) UpdateOnline(id int, domain domain.Online) (onlineClassObj domain.Online, err error) {
	var newRecord OnlineClass
	record := fromOnlineDomain(domain)
	err = cr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":                id,
		"trainer_id":        record.TrainerID,
		"class_category_id": record.ClassCategoryID,
		"date":              record.Date,
		"started_at":        record.StartedAt,
		"ended_at":          record.EndedAt,
		"link":              record.Link,
		"created_at":        record.CreatedAt,
		"updated_at":        record.UpdatedAt,
	}).Error

	if err != nil {
		return onlineClassObj, err
	}

	onlineClassObj = toOnlineDomain(newRecord)
	return onlineClassObj, nil
}

// Delete implements domain.Repository
func (cr classRepository) DeleteCategory(id int) (err error) {
	var record ClassCategory
	return cr.DB.Delete(&record, id).Error
}

// Update implements domain.Repository
func (cr classRepository) UpdateCategory(id int, domain domain.Category) (categoryObj domain.Category, err error) {
	var newRecord ClassCategory
	record := fromCategoryDomain(domain)
	err = cr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":           id,
		"name":         record.Name,
		"description":  record.Description,
		"picture_link": record.PictureLink,
		"created_at":   record.CreatedAt,
		"updated_at":   record.UpdatedAt,
	}).Error

	if err != nil {
		return categoryObj, err
	}

	categoryObj = toCategoryDomain(newRecord)
	return categoryObj, nil
}

// GetByID implements domain.Repository
func (cr classRepository) GetCategoryByID(id int) (domain domain.Category, err error) {
	var newRecord ClassCategory
	err = cr.DB.First(&newRecord, id).Error

	if err != nil {
		return domain, err
	}

	return toCategoryDomain(newRecord), nil
}

// Get implements domain.Repository
func (cr classRepository) GetCategory() (categoryObj []domain.Category, err error) {
	var newRecords []ClassCategory

	err = cr.DB.Find(&newRecords).Error

	if err != nil {
		return categoryObj, err
	}

	for _, value := range newRecords {
		categoryObj = append(categoryObj, toCategoryDomain(value))
	}

	return categoryObj, nil
}

// GetByEmail implements domain.Repository
func (cr classRepository) GetByEmail(email string) (categoryObj domain.Category, err error) {
	var newRecord ClassCategory
	err = cr.DB.Where("email = ?", email).First(&newRecord).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return categoryObj, err
	}

	return toCategoryDomain(newRecord), nil
}

// Create implements domain.Repository
func (cr classRepository) CreateCategory(domain domain.Category) (categoryObj domain.Category, err error) {
	// var recordDetail ClassDetail
	newRecord := fromCategoryDomain(domain)
	err = cr.DB.Create(&newRecord).Error

	if err != nil {
		return categoryObj, err
	}

	categoryObj = toCategoryDomain(newRecord)
	return categoryObj, nil
}

func NewClassRepository(db *gorm.DB) domain.Repository {
	return classRepository{
		DB: db,
	}
}
