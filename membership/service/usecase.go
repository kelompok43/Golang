package serviceMembership

import (
	"fmt"
	"strconv"
	"time"

	timeHelper "github.com/kelompok43/Golang/helpers/time"
	"github.com/kelompok43/Golang/membership/domain"
)

type membershipService struct {
	repository domain.Repository
}

// GetAllData implements domain.Service
func (ms membershipService) GetAllData() (membershipObj []domain.Membership, err error) {
	membershipObj, _ = ms.repository.Get()

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

	timeNow, _ := strconv.Atoi(timeHelper.Timestamp())
	// expired := timeNow * category.Duration

	if category.Duration == 30 {
		domain.ExpiredAt = strconv.Itoa(int((time.Now().UnixNano() / 1000000) + 2629743000))
	}

	/*Human-readable time 	Seconds
	1 hour	3600000 seconds
	1 day	86400000 seconds
	1 week	604800000 seconds
	1 month (30.44 days) 	2.629.743.000 seconds
	2592000000
	2628000000
	1 year (365.24 days) 	 31556926000 seconds */

	domain.UserID = userID
	domain.MembershipCategoryID = categoryID
	domain.CreatedAt = strconv.Itoa(timeNow)
	domain.UpdatedAt = strconv.Itoa(timeNow)
	isFound, _ := ms.repository.GetByUserID(userID)

	if isFound.UserID == userID {
		membership, err := ms.repository.GetByID(isFound.ID)

		if err != nil {
			return membershipObj, err
		}

		domain.ExpiredAt = strconv.Itoa(int((time.Now().UnixNano() / 1000000) + 2629743000))
		fmt.Println("expired  at update", domain.ExpiredAt)
		domain.CreatedAt = membership.CreatedAt
		domain.ID = membership.ID
		membershipObj, err = ms.repository.Update(membership.ID, domain)

		if err != nil {
			return membershipObj, err
		}

		return membershipObj, nil
	}

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
