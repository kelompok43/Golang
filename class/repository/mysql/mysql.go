package repoMySQL

import (
	"errors"

	"github.com/kelompok43/Golang/class/domain"
	"gorm.io/gorm"
)

type classRepository struct {
	DB *gorm.DB
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
func (classRepository) DeleteOnline(id int) (err error) {
	panic("unimplemented")
}

// GetCaOnlineID implements domain.Repository
func (classRepository) GetOnlineByID(id int) (categoryObj domain.Online, err error) {
	panic("unimplemented")
}

// GetOnline implements domain.Repository
func (cr classRepository) GetOnline() (categoryObj []domain.Online, err error) {
	var newRecords []OnlineClass

	err = cr.DB.Find(&newRecords).Error

	if err != nil {
		return categoryObj, err
	}

	for _, value := range newRecords {
		categoryObj = append(categoryObj, toOnlineDomain(value))
	}

	return categoryObj, nil

}

// UpdateOnline implements domain.Repository
func (classRepository) UpdateOnline(id int, domain domain.Online) (categoryObj domain.Online, err error) {
	panic("unimplemented")
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
