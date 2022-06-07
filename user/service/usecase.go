package service

import "github.com/kelompok43/Golang/user/domain"

type userService struct {
	repository domain.Repository
}

func (us userService) CreateToken(email, password string) (token string, err error) {
	return
}

func (us userService) InsertData(domain domain.User) (userObj domain.User, err error) {
	return
}

// GetByEmailPassword implements domain.Service
func (userService) GetByEmailPassword(email string, password string) (id int, status string, err error) {
	panic("unimplemented")
}

func NewUserService(repo domain.Repository) domain.Service {
	return userService{
		repository: repo,
	}
}
