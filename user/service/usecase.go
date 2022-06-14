package service

import (
	"errors"

	authMiddleware "github.com/kelompok43/Golang/auth/middlewares"
	"github.com/kelompok43/Golang/helpers/encrypt"
	encryptHelper "github.com/kelompok43/Golang/helpers/encrypt"
	timeHelper "github.com/kelompok43/Golang/helpers/time"
	"github.com/kelompok43/Golang/user/domain"
)

type userService struct {
	repository domain.Repository
	jwtAuth    authMiddleware.ConfigJWT
}

// ChangePassword implements domain.Service
func (us userService) ChangePassword(id int, domain domain.User) (userObj domain.User, err error) {
	domain.ID = id
	domain.UpdatedAt = timeHelper.Timestamp()
	userObj, err = us.repository.Update(domain)

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

// GetByEmail implements domain.Service
func (us userService) GetByEmail(email string) (userObj domain.User, err error) {
	userObj, err = us.repository.GetByEmail(email)

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

// InsertDetailData implements domain.Service
func (us userService) InsertDetailData(domain domain.User) (userObj domain.User, err error) {
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	userObj, err = us.repository.AddDetail(domain)

	if err != nil {
		return userObj, err
	}

	userObj, err = us.GetByID(userObj.ID)

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

// GetByID implements domain.Service
func (us userService) GetByID(id int) (userObj domain.User, err error) {
	userObj, err = us.repository.GetByID(id)

	if err != nil {
		return userObj, err
	}

	detail, _ := us.repository.GetDetail(id)

	userObj.DOB = detail.DOB
	userObj.Phone = detail.Phone
	userObj.Address = detail.Address
	userObj.Gender = detail.Gender
	return userObj, nil
}

func (us userService) CreateToken(email, password string) (token string, userObj domain.User, err error) {
	userObj, err = us.repository.GetByEmail(email)

	if err != nil {
		return token, userObj, err
	}

	if !encrypt.ValidateHash(password, userObj.Password) {
		return token, userObj, errors.New("email atau kata sandi salah")
	}

	id := userObj.ID
	token, err = us.jwtAuth.GenerateToken(id)

	if err != nil {
		return token, userObj, err
	}

	userObj, err = us.GetByID(id)

	if err != nil {
		return token, userObj, err
	}

	return token, userObj, nil
}

func (us userService) InsertData(domain domain.User) (userObj domain.User, err error) {
	email := domain.Email
	_, errGetUser := us.repository.GetByEmail(email)

	if errGetUser == nil {
		return userObj, errors.New("email telah terdaftar")
	}

	domain.Password, err = encryptHelper.Hash(domain.Password)

	if err != nil {
		return userObj, err
	}

	domain.Status = "Bukan Member"
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	userObj, err = us.repository.Create(domain)

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

// GetAllData implements domain.Service
func (us userService) GetAllData() (userObj []domain.User, err error) {
	userObj, _ = us.repository.Get()

	if err != nil {
		return userObj, err
	}

	return userObj, nil
}

func NewUserService(repo domain.Repository, jwtAuth authMiddleware.ConfigJWT) domain.Service {
	return userService{
		repository: repo,
		jwtAuth:    jwtAuth,
	}
}
