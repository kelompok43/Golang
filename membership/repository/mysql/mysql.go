package repoMySQL

import (
	"github.com/kelompok43/Golang/membership/domain"
	"gorm.io/gorm"
)

type membershipRepository struct {
	DB *gorm.DB
}

// Create implements domain.Repository
func (mr membershipRepository) Create(domain domain.Membership) (membershipObj domain.Membership, err error) {
	newRecord := fromDomain(domain)
	err = mr.DB.Create(&newRecord).Error

	if err != nil {
		return membershipObj, err
	}

	membershipObj = toDomain(newRecord)
	return membershipObj, nil
}

// GetByUserID implements domain.Repository
func (mr membershipRepository) GetByUserID(userID int) (membershipObj domain.Membership, err error) {
	var newRecord Membership
	err = mr.DB.Where("user_id = ?", userID).Find(&newRecord).Error

	if err != nil {
		return membershipObj, err
	}

	membershipObj = toDomain(newRecord)

	return membershipObj, nil
}

// GetByPrice implements domain.Repository
func (mr membershipRepository) GetCategoryByPrice(price int) (membershipCategoryObj domain.MembershipCategory, err error) {
	var newRecord MembershipCategory
	err = mr.DB.Where("price = ?", price).First(&newRecord).Error

	if err != nil {
		return membershipCategoryObj, err
	}

	membershipCategoryObj = categoryToDomain(newRecord)
	return membershipCategoryObj, nil
}

// CreateOrder implements domain.Repository
func (mr membershipRepository) CreateOrder(domain domain.MembershipOrder) (membershipOrderObj domain.MembershipOrder, err error) {
	newRecord := fromDomainToOrder(domain)
	err = mr.DB.Create(&newRecord).Error

	if err != nil {
		return membershipOrderObj, err
	}

	membershipOrderObj = orderToDomain(newRecord)
	return membershipOrderObj, nil
}

// Delete implements domain.Repository
func (mr membershipRepository) DeleteCategory(id int) (err error) {
	var record MembershipCategory
	return mr.DB.Delete(&record, id).Error
}

// Update implements domain.Repository
func (mr membershipRepository) UpdateCategory(id int, domain domain.MembershipCategory) (membershipCategoryObj domain.MembershipCategory, err error) {
	var newRecord MembershipCategory
	record := fromDomainToCategory(domain)
	err = mr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":         id,
		"category":   record.Category,
		"price":      record.Price,
		"duration":   record.Duration,
		"created_at": record.CreatedAt,
		"updated_at": record.UpdatedAt,
	}).Error

	if err != nil {
		return membershipCategoryObj, err
	}

	membershipCategoryObj = categoryToDomain(newRecord)
	return membershipCategoryObj, nil
}

// GetByID implements domain.Repository
func (mr membershipRepository) GetCategoryByID(id int) (domain domain.MembershipCategory, err error) {
	var newRecord MembershipCategory
	err = mr.DB.First(&newRecord, id).Error

	if err != nil {
		return domain, err
	}

	return categoryToDomain(newRecord), nil
}

// Get implements domain.Repository
func (mr membershipRepository) GetCategory() (membershipCategoryObj []domain.MembershipCategory, err error) {
	var newRecords []MembershipCategory

	err = mr.DB.Find(&newRecords).Error

	if err != nil {
		return membershipCategoryObj, err
	}

	for _, value := range newRecords {
		membershipCategoryObj = append(membershipCategoryObj, categoryToDomain(value))
	}

	return membershipCategoryObj, nil
}

// Create implements domain.Repository
func (mr membershipRepository) CreateCategory(domain domain.MembershipCategory) (membershipCategoryObj domain.MembershipCategory, err error) {
	// var recordDetail MembershipDetail
	newRecord := fromDomainToCategory(domain)
	err = mr.DB.Create(&newRecord).Error

	if err != nil {
		return membershipCategoryObj, err
	}

	membershipCategoryObj = categoryToDomain(newRecord)
	return membershipCategoryObj, nil
}

func NewMembershipRepository(db *gorm.DB) domain.Repository {
	return membershipRepository{
		DB: db,
	}
}
