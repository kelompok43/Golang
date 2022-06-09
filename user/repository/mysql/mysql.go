package repoMySQL

import (
	"errors"
	"fmt"

	"github.com/kelompok43/Golang/user/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
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

	ur.DB.Model(&User{}).Select("*").Joins("left join user_details on user_details.user_id = users.id").Scan(&newRecords)

	fmt.Println(err)

	// if err != nil {
	// 	return userObj, err
	// }

	for _, value := range newRecords {
		userObj = append(userObj, toDomain(value))
	}

	fmt.Println(userObj)
	return userObj, nil
}

// GetByEmail implements domain.Repository
func (ur userRepository) GetByEmail(email string) (userObj domain.User, err error) {
	var newRecord User
	err = ur.DB.Where("email = ?", email).First(&newRecord).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userObj, err
	}

	return userObj, nil
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

// GetByEmailPassword implements domain.Repository
func (ur userRepository) GetByEmailPassword(email string, password string) (domain domain.User, err error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return userRepository{
		DB: db,
	}
}
