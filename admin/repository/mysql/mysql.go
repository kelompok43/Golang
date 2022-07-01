package repoMySQL

import (
	"errors"

	"github.com/kelompok43/Golang/admin/domain"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

// Update implements domain.Repository
func (ur adminRepository) Update(domain domain.Admin) (adminObj domain.Admin, err error) {
	var newRecord Admin
	rec := fromDomain(domain)
	err = ur.DB.Model(&newRecord).Where("id = ?", domain.ID).Updates(map[string]interface{}{
		"id":         rec.ID,
		"name":       rec.Name,
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
func (ur adminRepository) GetByID(id int) (domain domain.Admin, err error) {
	var newRecord Admin
	err = ur.DB.First(&newRecord, id).Error

	if err != nil {
		return domain, err
	}

	return toDomain(newRecord), nil
}

// Get implements domain.Repository
func (ur adminRepository) Get() (adminObj []domain.Admin, err error) {
	var newRecords []Admin

	err = ur.DB.Find(&newRecords).Error

	if err != nil {
		return adminObj, err
	}

	for _, value := range newRecords {
		adminObj = append(adminObj, toDomain(value))
	}

	return adminObj, nil
}

// GetByEmail implements domain.Repository
func (ur adminRepository) GetByEmail(email string) (adminObj domain.Admin, err error) {
	var newRecord Admin
	err = ur.DB.Where("email = ?", email).First(&newRecord).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return adminObj, err
	}

	return toDomain(newRecord), nil
}

// Create implements domain.Repository
func (ur adminRepository) Create(domain domain.Admin) (adminObj domain.Admin, err error) {
	// var recordDetail AdminDetail
	newRecord := fromDomain(domain)
	err = ur.DB.Create(&newRecord).Error

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
