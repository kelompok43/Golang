package repoMySQL

import (
	"errors"

	"github.com/kelompok43/Golang/news/domain"
	"gorm.io/gorm"
)

type newsRepository struct {
	DB *gorm.DB
}

// CreateCategory implements domain.Repository
func (nr newsRepository) CreateCategory(domain domain.Category) (newsCategoryObj domain.Category, err error) {
	newRecord := fromCategoryDomain(domain)
	err = nr.DB.Create(&newRecord).Error

	if err != nil {
		return newsCategoryObj, err
	}

	newsCategoryObj = toCategoryDomain(newRecord)
	return newsCategoryObj, nil
}

// DeleteCategory implements domain.Repository
func (nr newsRepository) DeleteCategory(id int) (err error) {
	var record NewsCategory
	return nr.DB.Delete(&record, id).Error
}

// GetCategory implements domain.Repository
func (nr newsRepository) GetCategory() (newsCategoryObj []domain.Category, err error) {
	var newRecords []NewsCategory

	err = nr.DB.Find(&newRecords).Error

	if err != nil {
		return newsCategoryObj, err
	}

	for _, value := range newRecords {
		newsCategoryObj = append(newsCategoryObj, toCategoryDomain(value))
	}

	return newsCategoryObj, nil
}

// GetCategoryByID implements domain.Repository
func (nr newsRepository) GetCategoryByID(id int) (newsCategoryObj domain.Category, err error) {
	var newRecord NewsCategory
	err = nr.DB.First(&newRecord, id).Error

	if err != nil {
		return newsCategoryObj, err
	}

	return toCategoryDomain(newRecord), nil
}

// UpdateCategory implements domain.Repository
func (nr newsRepository) UpdateCategory(id int, domain domain.Category) (newsCategoryObj domain.Category, err error) {
	var newRecord NewsCategory
	record := fromCategoryDomain(domain)
	err = nr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":         id,
		"name":       record.Name,
		"created_at": record.CreatedAt,
		"updated_at": record.UpdatedAt,
	}).Error

	if err != nil {
		return newsCategoryObj, err
	}

	newsCategoryObj = toCategoryDomain(newRecord)
	return newsCategoryObj, nil
}

// Delete implements domain.Repository
func (nr newsRepository) Delete(id int) (err error) {
	var record News
	return nr.DB.Delete(&record, id).Error
}

// Update implements domain.Repository
func (nr newsRepository) Update(id int, domain domain.News) (newsObj domain.News, err error) {
	var newRecord News
	record := fromDomain(domain)
	err = nr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":               id,
		"news_category_id": record.NewsCategoryID,
		"title":            record.Title,
		"author":           record.Author,
		"date":             record.Date,
		"description":      record.Description,
		"picture_link":     record.PictureLink,
		"created_at":       record.CreatedAt,
		"updated_at":       record.UpdatedAt,
	}).Error

	if err != nil {
		return newsObj, err
	}

	newsObj = toDomain(newRecord)
	return newsObj, nil
}

// GetByID implements domain.Repository
func (nr newsRepository) GetByID(id int) (newsObj domain.News, err error) {
	var newRecord News
	err = nr.DB.First(&newRecord, id).Error

	if err != nil {
		return newsObj, err
	}

	return toDomain(newRecord), nil
}

// Get implements domain.Repository
func (nr newsRepository) Get() (newsObj []domain.News, err error) {
	var newRecords []News

	err = nr.DB.Find(&newRecords).Error

	if err != nil {
		return newsObj, err
	}

	for _, value := range newRecords {
		newsObj = append(newsObj, toDomain(value))
	}

	return newsObj, nil
}

// GetByEmail implements domain.Repository
func (nr newsRepository) GetByEmail(email string) (newsObj domain.News, err error) {
	var newRecord News
	err = nr.DB.Where("email = ?", email).First(&newRecord).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return newsObj, err
	}

	return toDomain(newRecord), nil
}

// Create implements domain.Repository
func (nr newsRepository) Create(domain domain.News) (newsObj domain.News, err error) {
	// var recordDetail NewsDetail
	newRecord := fromDomain(domain)
	err = nr.DB.Create(&newRecord).Error

	if err != nil {
		return newsObj, err
	}

	newsObj = toDomain(newRecord)
	return newsObj, nil
}

func NewNewsRepository(db *gorm.DB) domain.Repository {
	return newsRepository{
		DB: db,
	}
}
