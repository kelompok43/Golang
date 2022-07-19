package service_test

import (
	"errors"
	"testing"

	"github.com/kelompok43/Golang/admin/domain"
	"github.com/kelompok43/Golang/admin/domain/mocks"
	"github.com/kelompok43/Golang/admin/service"
	authMiddleware "github.com/kelompok43/Golang/auth/middlewares"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	repo mocks.Repository
	serv domain.Service
	dom  domain.Admin
)

func TestInsertData(t *testing.T) {
	serv = service.NewAdminService(&repo, authMiddleware.ConfigJWT{})
	dom = domain.Admin{
		ID:       1,
		Name:     "jiran",
		DOB:      "03-12-2000",
		Gender:   "L",
		Address:  "jl. jalan",
		Role:     "Super Admin",
		Email:    "jiran@mail.com",
		Password: "mypassword",
	}

	t.Run("Admin Service - Insert Data || Valid", func(t *testing.T) {
		repo.On("GetByEmail", mock.AnythingOfType("string")).Return(domain.Admin{}, errors.New("error")).Once()
		repo.On("Create", mock.AnythingOfType("domain.Admin")).Return(dom, nil).Once()
		adminObj, err := serv.InsertData(dom)

		assert.Nil(t, err)
		assert.Contains(t, adminObj.Email, "@mail.com")
	})
	t.Run("Admin Service - Insert Data || Invalid", func(t *testing.T) {
		repo.On("GetByEmail", mock.AnythingOfType("string")).Return(domain.Admin{}, errors.New("error")).Once()
		repo.On("Create", mock.AnythingOfType("domain.Admin")).Return(domain.Admin{}, errors.New("error")).Once()
		_, err := serv.InsertData(dom)

		assert.Error(t, err)
	})
}

func TestGetAllData(t *testing.T) {
	serv = service.NewAdminService(&repo, authMiddleware.ConfigJWT{})
	doms := []domain.Admin{
		{ID: 1,
			Name:     "jiran",
			DOB:      "03-12-2000",
			Gender:   "L",
			Address:  "jl. jalan",
			Role:     "Super Admin",
			Email:    "jiran@mail.com",
			Password: "mypassword",
		},
	}

	t.Run("Admin Service - Get All Data || Valid", func(t *testing.T) {
		repo.On("Get").Return(doms, nil).Once()
		adminObj, err := serv.GetAllData()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(adminObj))
	})
	t.Run("Admin Service - Get All Data || Invalid", func(t *testing.T) {
		repo.On("Get").Return([]domain.Admin{}, errors.New("error")).Once()
		_, err := serv.GetAllData()

		assert.Error(t, err)
	})
}

func TestGetDataByID(t *testing.T) {
	serv = service.NewAdminService(&repo, authMiddleware.ConfigJWT{})
	dom = domain.Admin{
		ID:       1,
		Name:     "jiran",
		DOB:      "03-12-2000",
		Gender:   "L",
		Address:  "jl. jalan",
		Role:     "Super Admin",
		Email:    "jiran@mail.com",
		Password: "mypassword",
	}

	t.Run("Admin Service - Get Data By ID || Valid", func(t *testing.T) {
		repo.On("GetByID", mock.AnythingOfType("int")).Return(dom, nil).Once()
		adminObj, err := serv.GetByID(dom.ID)

		assert.Nil(t, err)
		assert.Contains(t, adminObj.Email, "@mail.com")
	})
	t.Run("Admin Service - Get Data By ID || Invalid", func(t *testing.T) {
		repo.On("GetByID", mock.AnythingOfType("int")).Return(domain.Admin{}, errors.New("error")).Once()
		_, err := serv.GetByID(dom.ID)

		assert.Error(t, err)
	})
}

func TestGetDataByEmail(t *testing.T) {
	serv = service.NewAdminService(&repo, authMiddleware.ConfigJWT{})
	dom = domain.Admin{
		ID:       1,
		Name:     "jiran",
		DOB:      "03-12-2000",
		Gender:   "L",
		Address:  "jl. jalan",
		Role:     "Super Admin",
		Email:    "jiran@mail.com",
		Password: "mypassword",
	}

	t.Run("Admin Service - Get Data By Email || Valid", func(t *testing.T) {
		repo.On("GetByEmail", mock.AnythingOfType("string")).Return(dom, nil).Once()
		adminObj, err := serv.GetByEmail(dom.Email)

		assert.Nil(t, err)
		assert.Contains(t, adminObj.Email, "@mail.com")
	})
	t.Run("Admin Service - Get Data By Email || Invalid", func(t *testing.T) {
		repo.On("GetByEmail", mock.AnythingOfType("string")).Return(domain.Admin{}, errors.New("error")).Once()
		_, err := serv.GetByEmail(dom.Email)

		assert.Error(t, err)
	})
}

func TestUpdateData(t *testing.T) {
	serv = service.NewAdminService(&repo, authMiddleware.ConfigJWT{})
	dom = domain.Admin{
		ID:       1,
		Name:     "jiran",
		DOB:      "03-12-2000",
		Gender:   "L",
		Address:  "jl. jalan",
		Role:     "Super Admin",
		Email:    "jiran@mail.com",
		Password: "mypassword",
	}

	t.Run("Admin Service - Update Data || Valid", func(t *testing.T) {
		repo.On("GetByID", mock.AnythingOfType("int")).Return(dom, nil).Once()
		repo.On("Update", mock.AnythingOfType("domain.Admin")).Return(dom, nil).Once()
		adminObj, err := serv.UpdateData(dom.ID, dom)

		assert.Nil(t, err)
		assert.Contains(t, adminObj.Email, "@mail.com")
	})
	t.Run("Admin Service - Update Data || Invalid", func(t *testing.T) {
		repo.On("GetByID", mock.AnythingOfType("int")).Return(domain.Admin{}, errors.New("error")).Once()
		repo.On("Update", mock.AnythingOfType("domain.Admin")).Return(domain.Admin{}, errors.New("error")).Once()
		_, err := serv.UpdateData(dom.ID, dom)

		assert.Error(t, err)
	})
}

func TestDeleteData(t *testing.T) {
	serv = service.NewAdminService(&repo, authMiddleware.ConfigJWT{})
	dom = domain.Admin{
		ID:       1,
		Name:     "jiran",
		DOB:      "03-12-2000",
		Gender:   "L",
		Address:  "jl. jalan",
		Role:     "Super Admin",
		Email:    "jiran@mail.com",
		Password: "mypassword",
	}

	t.Run("Admin Service - Delete Data || Valid", func(t *testing.T) {
		repo.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()
		err := serv.DeleteData(dom.ID)

		assert.Nil(t, err)
	})
	t.Run("Admin Service - Delete Data || Invalid", func(t *testing.T) {
		repo.On("Delete", mock.AnythingOfType("int")).Return(errors.New("error")).Once()
		err := serv.DeleteData(dom.ID)

		assert.Error(t, err)
	})
}

// func TestChangePassword(t *testing.T) {
// 	serv = service.NewAdminService(&repo, authMiddleware.ConfigJWT{})
// 	dom = domain.Admin{
// 		ID:       1,
// 		Name:     "jiran",
// 		DOB:      "03-12-2000",
// 		Gender:   "L",
// 		Address:  "jl. jalan",
// 		Role:     "Super Admin",
// 		Email:    "jiran@mail.com",
// 		Password: "mypassword",
// 	}

// 	t.Run("Admin Service - Change Password || Valid", func(t *testing.T) {
// 		repo.On("GetByID", mock.AnythingOfType("int")).Return(dom, nil).Once()
// 		repo.On("Update", mock.AnythingOfType("domain.Admin")).Return(dom, nil).Once()
// 		adminObj, err := serv.ChangePassword(dom.ID, dom)

// 		assert.Nil(t, err)
// 		assert.Contains(t, adminObj.Email, "jiran")
// 	})
// 	t.Run("Admin Service - Change Password || Invalid", func(t *testing.T) {
// 		repo.On("GetByID", mock.AnythingOfType("int")).Return(domain.Admin{}, errors.New("error")).Once()
// 		repo.On("Update", mock.AnythingOfType("domain.Admin")).Return(domain.Admin{}, errors.New("error")).Once()
// 		_, err := serv.ChangePassword(dom.ID, dom)

// 		assert.Error(t, err)
// 	})
// }
