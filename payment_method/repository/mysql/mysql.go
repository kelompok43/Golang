package repoMySQL

import (
	"github.com/kelompok43/Golang/payment_method/domain"
	"gorm.io/gorm"
)

type paymentMethodRepository struct {
	DB *gorm.DB
}

// Delete implements domain.Repository
func (tr paymentMethodRepository) Delete(id int) (err error) {
	var record PaymentMethod
	return tr.DB.Delete(&record, id).Error
}

// Update implements domain.Repository
func (tr paymentMethodRepository) Update(id int, domain domain.PaymentMethod) (paymentMethodObj domain.PaymentMethod, err error) {
	var newRecord PaymentMethod
	record := fromDomain(domain)
	err = tr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":         id,
		"name":       record.Name,
		"acc_number": record.AccNumber,
		"acc_name":   record.AccName,
		"created_at": record.CreatedAt,
		"updated_at": record.UpdatedAt,
	}).Error

	if err != nil {
		return paymentMethodObj, err
	}

	paymentMethodObj = toDomain(newRecord)
	return paymentMethodObj, nil
}

// GetByID implements domain.Repository
func (tr paymentMethodRepository) GetByID(id int) (domain domain.PaymentMethod, err error) {
	var newRecord PaymentMethod
	err = tr.DB.First(&newRecord, id).Error

	if err != nil {
		return domain, err
	}

	return toDomain(newRecord), nil
}

// Get implements domain.Repository
func (tr paymentMethodRepository) Get() (paymentMethodObj []domain.PaymentMethod, err error) {
	var newRecords []PaymentMethod

	err = tr.DB.Find(&newRecords).Error

	if err != nil {
		return paymentMethodObj, err
	}

	for _, value := range newRecords {
		paymentMethodObj = append(paymentMethodObj, toDomain(value))
	}

	return paymentMethodObj, nil
}

// Create implements domain.Repository
func (tr paymentMethodRepository) Create(domain domain.PaymentMethod) (paymentMethodObj domain.PaymentMethod, err error) {
	// var recordDetail PaymentMethodDetail
	newRecord := fromDomain(domain)
	err = tr.DB.Create(&newRecord).Error

	if err != nil {
		return paymentMethodObj, err
	}

	paymentMethodObj = toDomain(newRecord)
	return paymentMethodObj, nil
}

func NewPaymentMethodRepository(db *gorm.DB) domain.Repository {
	return paymentMethodRepository{
		DB: db,
	}
}
