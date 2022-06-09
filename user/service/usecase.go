package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/kelompok43/Golang/user/domain"
)

type userService struct {
	repository domain.Repository
}

// GetByID implements domain.Service
func (us userService) GetByID(id int) (userObj domain.User, err error) {
	userObj, err = us.repository.GetByID(id)

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

func (us userService) CreateToken(email, password string) (token string, err error) {
	return
}

func (us userService) InsertData(domain domain.User) (userObj domain.User, err error) {
	email := domain.Email

	_, errGetUser := us.repository.GetByEmail(email)

	if errGetUser == nil {
		return userObj, errors.New("email telah terdaftar")
	}

	domain.Status = "Not Member"
	domain.CreatedAt = strconv.Itoa(int(time.Now().UnixNano() / 1000000))
	domain.UpdatedAt = strconv.Itoa(int(time.Now().UnixNano() / 1000000))
	userObj, err = us.repository.Create(domain)

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

// GetAllData implements domain.Service
func (us userService) GetAllData() (userObj []domain.User, err error) {
	userObj, _ = us.repository.Get()

	// if err != nil {
	// 	return userObj, err
	// }

	return userObj, nil
}

func (userService) GetByEmailPassword(email string, password string) (id int, status string, err error) {
	panic("unimplemented")
}

func NewUserService(repo domain.Repository) domain.Service {
	return userService{
		repository: repo,
	}
}
