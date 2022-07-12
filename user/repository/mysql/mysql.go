package repoMySQL

import (
	"errors"

	"github.com/kelompok43/Golang/user/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

// UpdateDetail implements domain.Repository
func (ur userRepository) UpdateDetail(domain domain.User) (userObj domain.User, err error) {
	newRecord := fromDomainToUserDetail(domain)
	err = ur.DB.Create(&newRecord).Error

	if err != nil {
		return userObj, err
	}

	user := joinResult{
		ID:          domain.ID,
		DOB:         domain.DOB,
		Phone:       domain.Phone,
		Address:     domain.Address,
		PictureLink: domain.PictureLink,
	}

	return toDomain(user), nil
}

// Update implements domain.Repository
func (ur userRepository) Update(domain domain.User) (userObj domain.User, err error) {
	var newRecord User
	var newDetailRecord UserDetail
	rec := fromDomainToUser(domain)
	err = ur.DB.Model(&newRecord).Where("id = ?", domain.ID).Updates(map[string]interface{}{
		"id":         rec.ID,
		"name":       rec.Name,
		"email":      rec.Email,
		"status":     rec.Status,
		"password":   rec.Password,
		"updated_at": domain.UpdatedAt,
	}).Error

	if err != nil {
		return userObj, err
	}

	detailRec := fromDomainToUserDetail(domain)
	err = ur.DB.Model(&newDetailRecord).Where("user_id = ?", domain.ID).Updates(map[string]interface{}{
		"user_id":      detailRec.UserID,
		"dob":          detailRec.DOB,
		"phone":        detailRec.Phone,
		"address":      detailRec.Address,
		"picture_link": detailRec.PictureLink,
		"updated_at":   domain.UpdatedAt,
	}).Error

	if err != nil {
		return userObj, err
	}

	user := joinResult{
		ID:       rec.ID,
		Name:     rec.Name,
		DOB:      detailRec.DOB,
		Email:    rec.Email,
		Password: rec.Password,
		Phone:    detailRec.Phone,
		Address:  detailRec.Address,
		// Gender:    detailRec.Gender,
		PictureLink: detailRec.PictureLink,
		Status:      rec.Status,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}

	return toDomain(user), nil
}

// GetDetail implements domain.Repository
func (ur userRepository) GetDetail(id int) (userObj domain.User, err error) {
	var record UserDetail
	err = ur.DB.Where("user_id = ?", id).First(&record).Error

	if err != nil {
		return userObj, err
	}

	user := joinResult{
		ID:          record.UserID,
		DOB:         record.DOB,
		Phone:       record.Phone,
		Address:     record.Address,
		PictureLink: record.PictureLink,
	}

	return toDomain(user), nil
}

// AddDetail implements domain.Repository
func (ur userRepository) AddDetail(domain domain.User) (userObj domain.User, err error) {
	newRecord := fromDomainToUserDetail(domain)
	err = ur.DB.Create(&newRecord).Error

	if err != nil {
		return userObj, err
	}

	user := joinResult{
		ID:          domain.ID,
		DOB:         domain.DOB,
		Phone:       domain.Phone,
		Address:     domain.Address,
		PictureLink: domain.PictureLink,
	}

	return toDomain(user), nil
}

// GetByID implements domain.Repository
func (ur userRepository) GetByID(id int) (domain domain.User, err error) {
	var newRecord User
	err = ur.DB.First(&newRecord, id).Error

	if err != nil {
		return domain, err
	}

	userObj := joinResult{
		ID:        newRecord.ID,
		Name:      newRecord.Name,
		Email:     newRecord.Email,
		Status:    newRecord.Status,
		CreatedAt: newRecord.CreatedAt,
		UpdatedAt: newRecord.UpdatedAt,
	}

	return toDomain(userObj), nil
}

// Get implements domain.Repository
func (ur userRepository) Get() (userObj []domain.User, err error) {
	var newRecords []joinResult

	err = ur.DB.Model(&UserDetail{}).Select("*").Joins("right join users on users.id = user_details.user_id").Scan(&newRecords).Error

	if err != nil {
		return userObj, err
	}

	for _, value := range newRecords {
		userObj = append(userObj, toDomain(value))
	}

	return userObj, nil
}

// GetByEmail implements domain.Repository
func (ur userRepository) GetByEmail(email string) (userObj domain.User, err error) {
	var newRecord User
	err = ur.DB.Where("email = ?", email).First(&newRecord).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userObj, err
	}

	userDetail, err := ur.GetDetail(newRecord.ID)

	if err != nil {
		return userObj, err
	}

	user := joinResult{
		ID:        newRecord.ID,
		Email:     newRecord.Email,
		Password:  newRecord.Password,
		Status:    newRecord.Status,
		DOB:       userDetail.DOB,
		Phone:     userDetail.Phone,
		Address:   userDetail.Address,
		CreatedAt: newRecord.CreatedAt,
		UpdatedAt: newRecord.UpdatedAt,
	}
	return toDomain(user), nil
}

// Create implements domain.Repository
func (ur userRepository) Create(domain domain.User) (userObj domain.User, err error) {
	// var recordDetail UserDetail
	record := fromDomainToUser(domain)
	err = ur.DB.Create(&record).Error

	if err != nil {
		return userObj, err
	}

	// userObj = toDomain(record, recordDetail)
	recordDetail := joinResult{
		ID:        record.ID,
		Name:      record.Name,
		Email:     record.Email,
		Password:  record.Password,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
	userObj = toDomain(recordDetail)
	return userObj, nil
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return userRepository{
		DB: db,
	}
}
