package repoMySQL

import (
	"errors"

	"github.com/kelompok43/Golang/admin/domain"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

// Delete implements domain.Repository
func (ar adminRepository) Delete(id int) (err error) {
	var record Admin
	return ar.DB.Delete(&record, id).Error
}

// Update implements domain.Repository
func (ar adminRepository) Update(domain domain.Admin) (adminObj domain.Admin, err error) {
	var newRecord Admin
	rec := fromDomain(domain)
	err = ar.DB.Model(&newRecord).Where("id = ?", domain.ID).Updates(map[string]interface{}{
		"id":         rec.ID,
		"name":       rec.Name,
		"dob":        rec.DOB,
		"gender":     rec.Gender,
		"address":    rec.Address,
		"role":       rec.Role,
		"email":      rec.Email,
		"password":   rec.Password,
		"updated_at": domain.UpdatedAt,
	}).Error

	if err != nil {
		return adminObj, err
	}

	return toDomain(newRecord), nil
}

// GetByID implements domain.Repository
func (ar adminRepository) GetByID(id int) (domain domain.Admin, err error) {
	var newRecord Admin
	err = ar.DB.First(&newRecord, id).Error

	if err != nil {
		return domain, err
	}

	return toDomain(newRecord), nil
}

// Get implements domain.Repository
func (ar adminRepository) Get() (adminObj []domain.Admin, err error) {
	var newRecords []Admin

	err = ar.DB.Find(&newRecords).Error

	if err != nil {
		return adminObj, err
	}

	for _, value := range newRecords {
		adminObj = append(adminObj, toDomain(value))
	}

	return adminObj, nil
}

// GetByEmail implements domain.Repository
func (ar adminRepository) GetByEmail(email string) (adminObj domain.Admin, err error) {
	var newRecord Admin
	err = ar.DB.Where("email = ?", email).First(&newRecord).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return adminObj, err
	}

	return toDomain(newRecord), nil
}

// Create implements domain.Repository
func (ar adminRepository) Create(domain domain.Admin) (adminObj domain.Admin, err error) {
	// var recordDetail AdminDetail
	newRecord := fromDomain(domain)
	err = ar.DB.Create(&newRecord).Error

	if err != nil {
		return adminObj, err
	}

	adminObj = toDomain(newRecord)
	return adminObj, nil
}

func NewAdminRepository(db *gorm.DB) domain.Repository {
	return adminRepository{
		DB: db,
	}
}
