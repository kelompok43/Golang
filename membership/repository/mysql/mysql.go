package repoMySQL

import (
	"github.com/kelompok43/Golang/membership/domain"
	"gorm.io/gorm"
)

type membershipRepository struct {
	DB *gorm.DB
}

// GetByUserID implements domain.Repository
func (mr membershipRepository) GetByUserID(userID int) (membershipOrderObj []domain.MembershipOrder, err error) {
	var newRecord []MembershipOrder
	err = mr.DB.Where("user_id = ?", userID).Find(&newRecord).Error

	if err != nil {
		return membershipOrderObj, err
	}

	for _, value := range newRecord {
		membershipOrderObj = append(membershipOrderObj, orderToDomain(value))
	}

	return membershipOrderObj, nil
}

// GetByPrice implements domain.Repository
func (mr membershipRepository) GetByPrice(price int) (membershipObj domain.Membership, err error) {
	var newRecord Membership
	err = mr.DB.Where("price = ?", price).First(&newRecord).Error

	if err != nil {
		return membershipObj, err
	}

	return membershipObj, nil
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
func (mr membershipRepository) Delete(id int) (err error) {
	var record Membership
	return mr.DB.Delete(&record, id).Error
}

// Update implements domain.Repository
func (mr membershipRepository) Update(id int, domain domain.Membership) (membershipObj domain.Membership, err error) {
	var newRecord Membership
	record := fromDomain(domain)
	err = mr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":         id,
		"category":   record.Category,
		"price":      record.Price,
		"duration":   record.Duration,
		"created_at": record.CreatedAt,
		"updated_at": record.UpdatedAt,
	}).Error

	if err != nil {
		return membershipObj, err
	}

	membershipObj = toDomain(newRecord)
	return membershipObj, nil
}

// GetByID implements domain.Repository
func (mr membershipRepository) GetByID(id int) (domain domain.Membership, err error) {
	var newRecord Membership
	err = mr.DB.First(&newRecord, id).Error

	if err != nil {
		return domain, err
	}

	return toDomain(newRecord), nil
}

// Get implements domain.Repository
func (mr membershipRepository) Get() (membershipObj []domain.Membership, err error) {
	var newRecords []Membership

	err = mr.DB.Find(&newRecords).Error

	if err != nil {
		return membershipObj, err
	}

	for _, value := range newRecords {
		membershipObj = append(membershipObj, toDomain(value))
	}

	return membershipObj, nil
}

// Create implements domain.Repository
func (mr membershipRepository) Create(domain domain.Membership) (membershipObj domain.Membership, err error) {
	// var recordDetail MembershipDetail
	newRecord := fromDomain(domain)
	err = mr.DB.Create(&newRecord).Error

	if err != nil {
		return membershipObj, err
	}

	membershipObj = toDomain(newRecord)
	return membershipObj, nil
}

func NewMembershipRepository(db *gorm.DB) domain.Repository {
	return membershipRepository{
		DB: db,
	}
}
