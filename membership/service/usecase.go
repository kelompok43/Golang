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

// GetByUserID implements domain.Service
func (ms membershipService) GetByUserID(userID int) (membershipObj domain.Membership, err error) {
	membershipObj, err = ms.repository.GetByUserID(userID)

	if err != nil {
		return membershipObj, err
	}

	return membershipObj, nil
}

// InsertData implements domain.Service
func (ms membershipService) InsertData(userID, categoryID int) (membershipObj domain.Membership, err error) {
	var domain domain.Membership
	category, err := ms.repository.GetCategoryByID(categoryID)

	if err != nil {
		return membershipObj, err
	}

	fmt.Println(userID)
	fmt.Println("category", category)
	timeNow, _ := strconv.Atoi(timeHelper.Timestamp())
	expired := timeNow * category.Duration
	domain.ExpiredAt = strconv.Itoa(expired)
	domain.UserID = userID
	domain.MembershipCategoryID = categoryID
	domain.CreatedAt = strconv.Itoa(timeNow)
	domain.UpdatedAt = strconv.Itoa(timeNow)
	membershipObj, err = ms.repository.Create(domain)

	if err != nil {
		return membershipObj, err
	}

	return membershipObj, nil
}

// GetCategoryByPrice implements domain.Service
func (ms membershipService) GetCategoryByPrice(price int) (membershipCategoryObj domain.MembershipCategory, err error) {
	membershipCategoryObj, err = ms.repository.GetCategoryByPrice(price)

	if err != nil {
		return membershipCategoryObj, err
	}

	return membershipCategoryObj, nil
}

// AddOrder implements domain.Service
func (ms membershipService) InsertOrder(transactionID, price int) (membershipOrderObj domain.MembershipOrder, err error) {
	var domain domain.MembershipOrder
	membership, err := ms.repository.GetCategoryByPrice(price)

	if err != nil {
		return membershipOrderObj, err
	}

	fmt.Println(price)
	fmt.Println(membership.ID)
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
func (ms membershipService) DeleteCategory(id int) (err error) {
	errResp := ms.repository.DeleteCategory(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// UpdateData implements domain.Service
func (ms membershipService) UpdateCategory(id int, domain domain.MembershipCategory) (membershipCategoryObj domain.MembershipCategory, err error) {
	membership, errGetByID := ms.GetCategoryByID(id)

	if errGetByID != nil {
		return membership, errGetByID
	}

	fmt.Println(membership)

	domain.CreatedAt = membership.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	membershipCategoryObj, err = ms.repository.UpdateCategory(id, domain)

	if err != nil {
		return membershipCategoryObj, err
	}

	return membershipCategoryObj, nil
}

// GetByID implements domain.Service
func (ms membershipService) GetCategoryByID(id int) (membershipCategoryObj domain.MembershipCategory, err error) {
	membershipCategoryObj, err = ms.repository.GetCategoryByID(id)

	if err != nil {
		return membershipCategoryObj, err
	}

	return membershipCategoryObj, nil
}

func (ms membershipService) InsertCategory(domain domain.MembershipCategory) (membershipCategoryObj domain.MembershipCategory, err error) {
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	membershipCategoryObj, err = ms.repository.CreateCategory(domain)

	if err != nil {
		return membershipCategoryObj, err
	}

	return membershipCategoryObj, nil
}

// GetAllData implements domain.Service
func (ms membershipService) GetAllCategory() (membershipCategoryObj []domain.MembershipCategory, err error) {
	membershipCategoryObj, _ = ms.repository.GetCategory()

	if err != nil {
		return membershipCategoryObj, err
	}

	return membershipCategoryObj, nil
}

func NewMembershipService(repo domain.Repository) domain.Service {
	return membershipService{
		repository: repo,
	}
}
