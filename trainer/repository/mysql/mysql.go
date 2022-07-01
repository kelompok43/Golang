package repoMySQL

import (
	"errors"

	"github.com/kelompok43/Golang/trainer/domain"
	"gorm.io/gorm"
)

type trainerRepository struct {
	DB *gorm.DB
}

// Delete implements domain.Repository
func (tr trainerRepository) Delete(id int) (err error) {
	var record Trainer
	return tr.DB.Delete(&record, id).Error
}

// Update implements domain.Repository
func (tr trainerRepository) Update(id int, domain domain.Trainer) (trainerObj domain.Trainer, err error) {
	var newRecord Trainer
	record := fromDomain(domain)
	err = tr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":           id,
		"name":         record.Name,
		"email":        record.Email,
		"dob":          record.DOB,
		"gender":       record.Gender,
		"phone":        record.Phone,
		"address":      record.Address,
		"picture_link": record.PictureLink,
		"field":        record.Field,
		"created_at":   record.CreatedAt,
		"updated_at":   record.UpdatedAt,
	}).Error

	if err != nil {
		return trainerObj, err
	}

	trainerObj = toDomain(newRecord)
	return trainerObj, nil
}

// GetByID implements domain.Repository
func (tr trainerRepository) GetByID(id int) (domain domain.Trainer, err error) {
	var newRecord Trainer
	err = tr.DB.First(&newRecord, id).Error

	if err != nil {
		return domain, err
	}

	return toDomain(newRecord), nil
}

// Get implements domain.Repository
func (tr trainerRepository) Get() (trainerObj []domain.Trainer, err error) {
	var newRecords []Trainer

	err = tr.DB.Find(&newRecords).Error

	if err != nil {
		return trainerObj, err
	}

	for _, value := range newRecords {
		trainerObj = append(trainerObj, toDomain(value))
	}

	return trainerObj, nil
}

// GetByEmail implements domain.Repository
func (tr trainerRepository) GetByEmail(email string) (trainerObj domain.Trainer, err error) {
	var newRecord Trainer
	err = tr.DB.Where("email = ?", email).First(&newRecord).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return trainerObj, err
	}

	return toDomain(newRecord), nil
}

// Create implements domain.Repository
func (tr trainerRepository) Create(domain domain.Trainer) (trainerObj domain.Trainer, err error) {
	// var recordDetail TrainerDetail
	newRecord := fromDomain(domain)
	err = tr.DB.Create(&newRecord).Error

	if err != nil {
		return trainerObj, err
	}

	trainerObj = toDomain(newRecord)
	return trainerObj, nil
}

func NewTrainerRepository(db *gorm.DB) domain.Repository {
	return trainerRepository{
		DB: db,
	}
}
