package service

import (
	"errors"

	"github.com/kelompok43/Golang/admin/domain"
	authMiddleware "github.com/kelompok43/Golang/auth/middlewares"
	encryptHelper "github.com/kelompok43/Golang/helpers/encrypt"
	timeHelper "github.com/kelompok43/Golang/helpers/time"
)

type adminService struct {
	repository domain.Repository
	jwtAuth    authMiddleware.ConfigJWT
}

// DeleteData implements domain.Service
func (as adminService) DeleteData(id int) (err error) {
	errResp := as.repository.Delete(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// UpdateData implements domain.Service
func (as adminService) UpdateData(id int, domain domain.Admin) (adminObj domain.Admin, err error) {
	admin, err := as.GetByID(id)

	if err != nil {
		return adminObj, err
	}

	domain.ID = admin.ID
	// domain.Password = admin.Password
	domain.CreatedAt = admin.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()

	adminObj, err = as.repository.Update(domain)

	if err != nil {
		return adminObj, err
	}

	return adminObj, nil
}

// ChangePassword implements domain.Service
func (as adminService) ChangePassword(id int, domain domain.Admin) (adminObj domain.Admin, err error) {
	admin, err := as.repository.GetByID(id)

	if err != nil {
		return adminObj, err
	}

	admin.ID = id
	admin.Password, err = encryptHelper.Hash(domain.Password)

	if err != nil {
		return adminObj, err
	}

	admin.UpdatedAt = timeHelper.Timestamp()
	adminObj, err = as.repository.Update(admin)

	if err != nil {
		return adminObj, err
	}

	return adminObj, nil
}

// GetByEmail implements domain.Service
func (as adminService) GetByEmail(email string) (adminObj domain.Admin, err error) {
	adminObj, err = as.repository.GetByEmail(email)

	if err != nil {
		return adminObj, err
	}

	return adminObj, nil
}

// GetByID implements domain.Service
func (as adminService) GetByID(id int) (adminObj domain.Admin, err error) {
	adminObj, err = as.repository.GetByID(id)

	if err != nil {
		return adminObj, err
	}

	return adminObj, nil
}

func (as adminService) CreateToken(email, password string) (token string, adminObj domain.Admin, err error) {
	adminObj, err = as.repository.GetByEmail(email)

	if err != nil {
		return token, adminObj, err
	}

	if !encryptHelper.ValidateHash(password, adminObj.Password) {
		return token, adminObj, errors.New("email atau kata sandi salah")
	}

	id := adminObj.ID
	token, err = as.jwtAuth.GenerateToken(id)

	if err != nil {
		return token, adminObj, err
	}

	adminObj, err = as.GetByID(id)

	if err != nil {
		return token, adminObj, err
	}

	return token, adminObj, nil
}

func (as adminService) InsertData(domain domain.Admin) (adminObj domain.Admin, err error) {
	email := domain.Email
	_, errGetAdmin := as.repository.GetByEmail(email)

	if errGetAdmin == nil {
		return adminObj, errors.New("email telah terdaftar")
	}

	domain.Password, err = encryptHelper.Hash(domain.Password)

	if err != nil {
		return adminObj, err
	}

	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	adminObj, err = as.repository.Create(domain)

	if err != nil {
		return adminObj, err
	}

	return adminObj, nil
}

// GetAllData implements domain.Service
func (as adminService) GetAllData() (adminObj []domain.Admin, err error) {
	adminObj, _ = as.repository.Get()

	if err != nil {
		return adminObj, err
	}

	return adminObj, nil
}

func NewAdminService(repo domain.Repository, jwtAuth authMiddleware.ConfigJWT) domain.Service {
	return adminService{
		repository: repo,
		jwtAuth:    jwtAuth,
	}
}
