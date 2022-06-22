package servicePaymentMethod

import (
	"fmt"

	timeHelper "github.com/kelompok43/Golang/helpers/time"
	"github.com/kelompok43/Golang/membership/domain"
)

type membershipService struct {
	repository domain.Repository
}

// DeleteData implements domain.Service
func (ms membershipService) DeleteData(id int) (err error) {
	errResp := ms.repository.Delete(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// UpdateData implements domain.Service
func (ms membershipService) UpdateData(id int, domain domain.Membership) (membershipObj domain.Membership, err error) {
	membership, errGetByID := ms.GetByID(id)

	if errGetByID != nil {
		return membership, errGetByID
	}

	fmt.Println(membership)

	domain.CreatedAt = membership.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	membershipObj, err = ms.repository.Update(id, domain)

	if err != nil {
		return membershipObj, err
	}

	return membershipObj, nil
}

// GetByID implements domain.Service
func (ms membershipService) GetByID(id int) (membershipObj domain.Membership, err error) {
	membershipObj, err = ms.repository.GetByID(id)

	if err != nil {
		return membershipObj, err
	}

	return membershipObj, nil
}

func (ms membershipService) InsertData(domain domain.Membership) (membershipObj domain.Membership, err error) {
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	membershipObj, err = ms.repository.Create(domain)

	if err != nil {
		return membershipObj, err
	}

	return membershipObj, nil
}

// GetAllData implements domain.Service
func (ms membershipService) GetAllData() (membershipObj []domain.Membership, err error) {
	membershipObj, _ = ms.repository.Get()

	if err != nil {
		return membershipObj, err
	}

	return membershipObj, nil
}

func NewPaymentMethodService(repo domain.Repository) domain.Service {
	return membershipService{
		repository: repo,
	}
}
