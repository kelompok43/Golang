package service

import (
	"errors"
	"fmt"
	"strconv"

	authMiddleware "github.com/kelompok43/Golang/auth/middlewares"
	encryptHelper "github.com/kelompok43/Golang/helpers/encrypt"
	timeHelper "github.com/kelompok43/Golang/helpers/time"
	membershipDomain "github.com/kelompok43/Golang/membership/domain"
	"github.com/kelompok43/Golang/user/domain"
)

type userService struct {
	repository        domain.Repository
	jwtAuth           authMiddleware.ConfigJWT
	membershipService membershipDomain.Service
}

// UpdateStatus implements domain.Service
func (us userService) UpdateStatus(id int) (userObj domain.User, err error) {
	user, err := us.repository.GetByID(id)
	fmt.Println("id user update ", id)

	if err != nil {
		return userObj, err
	}

	// status := user.Status
	now, _ := strconv.Atoi(timeHelper.Timestamp())
	userMember, _ := us.membershipService.GetByUserID(id)
	userExpired, _ := strconv.Atoi(userMember.ExpiredAt)

	fmt.Println("user expired = ", userExpired)
	fmt.Println("now = ", now)
	fmt.Println("status = ", user.Status)
	if now < userExpired {
		user.Status = "Member"
	}

	fmt.Println("status2 = ", user.Status)
	// user.Status = status
	user.UpdatedAt = strconv.Itoa(now)
	userObj, err = us.repository.Update(user)

	if err != nil {
		return userObj, err
	}

	return userObj, nil
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

	if !encryptHelper.ValidateHash(password, userObj.Password) {
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

func NewUserService(repo domain.Repository, jwtAuth authMiddleware.ConfigJWT, ms membershipDomain.Service) domain.Service {
	return userService{
		repository:        repo,
		jwtAuth:           jwtAuth,
		membershipService: ms,
	}
}
