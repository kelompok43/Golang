package repoMySQL

import (
	"errors"

	"github.com/kelompok43/Golang/transaction/domain"
	"gorm.io/gorm"
)

type transactionRepository struct {
	DB *gorm.DB
}

// CreateDetail implements domain.Repository
func (tr transactionRepository) CreateDetail(domain domain.TransactionDetail) (transactionDetailObj domain.TransactionDetail, err error) {
	newRecord := fromDomainToDetail(domain)

	err = tr.DB.Create(&newRecord).Error

	if err != nil {
		return transactionDetailObj, err
	}

	transactionDetailObj = detailToDomain(newRecord)
	return transactionDetailObj, nil
}

// Delete implements domain.Repository
func (tr transactionRepository) Delete(id int) (err error) {
	var record Transaction
	return tr.DB.Delete(&record, id).Error
}

// Update implements domain.Repository
func (tr transactionRepository) Update(id int, domain domain.Transaction) (transactionObj domain.Transaction, err error) {
	var newRecord Transaction
	record := fromDomain(domain)
	err = tr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":                id,
		"user_id":           record.UserID,
		"payment_method_id": record.PaymentMethodID,
		"status":            record.Status,
		"total_price":       record.TotalPrice,
		"picture_link":      record.PictureLink,
		"created_at":        record.CreatedAt,
		"updated_at":        record.UpdatedAt,
	}).Error

	if err != nil {
		return transactionObj, err
	}

	transactionObj = toDomain(newRecord)
	return transactionObj, nil
}

// GetByID implements domain.Repository
func (tr transactionRepository) GetByID(id int) (domain domain.Transaction, err error) {
	var newRecord Transaction
	err = tr.DB.First(&newRecord, id).Error

	if err != nil {
		return domain, err
	}

	return toDomain(newRecord), nil
}

// Get implements domain.Repository
func (tr transactionRepository) Get() (transactionObj []domain.Transaction, err error) {
	var newRecords []Transaction

	err = tr.DB.Find(&newRecords).Error

	if err != nil {
		return transactionObj, err
	}

	for _, value := range newRecords {
		transactionObj = append(transactionObj, toDomain(value))
	}

	return transactionObj, nil
}

// GetByEmail implements domain.Repository
func (tr transactionRepository) GetByEmail(email string) (transactionObj domain.Transaction, err error) {
	var newRecord Transaction
	err = tr.DB.Where("email = ?", email).First(&newRecord).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return transactionObj, err
	}

	return toDomain(newRecord), nil
}

// Create implements domain.Repository
func (tr transactionRepository) Create(domain domain.Transaction) (transactionObj domain.Transaction, err error) {
	// var recordDetail TransactionDetail
	newRecord := fromDomain(domain)
	err = tr.DB.Create(&newRecord).Error

	if err != nil {
		return transactionObj, err
	}

	transactionObj = toDomain(newRecord)
	return transactionObj, nil
}

func NewTransactionRepository(db *gorm.DB) domain.Repository {
	return transactionRepository{
		DB: db,
	}
}
