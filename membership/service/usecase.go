package serviceMembership

import (
	"fmt"
	"strconv"

	timeHelper "github.com/kelompok43/Golang/helpers/time"
	"github.com/kelompok43/Golang/membership/domain"
)

type membershipService struct {
	repository domain.Repository
}

// GetOrderByUserID implements domain.Service
func (ms membershipService) GetOrderByUserID(userID int) (membershipOrderObj []domain.MembershipOrder, err error) {
	membershipOrderObj, err = ms.repository.GetByUserID(userID)

	if err != nil {
		return membershipOrderObj, err
	}

	return membershipOrderObj, nil
}

// AddOrder implements domain.Service
func (ms membershipService) InsertOrder(transactionID, price int) (membershipOrderObj domain.MembershipOrder, err error) {
	var domain domain.MembershipOrder
	membership, err := ms.repository.GetByPrice(price)

	if err != nil {
		return membershipOrderObj, err
	}

	timeNow, _ := strconv.Atoi(timeHelper.Timestamp())
	expired := timeNow * membership.Duration
	domain.Expired = strconv.Itoa(expired)
	domain.MembershipID = membership.ID
	domain.TransactionID = transactionID
	domain.CreatedAt = strconv.Itoa(timeNow)
	domain.UpdatedAt = strconv.Itoa(timeNow)
	membershipOrderObj, err = ms.repository.CreateOrder(domain)

	if err != nil {
		return membershipOrderObj, err
	}

	return membershipOrderObj, nil
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

func NewMembershipService(repo domain.Repository) domain.Service {
	return membershipService{
		repository: repo,
	}
}
