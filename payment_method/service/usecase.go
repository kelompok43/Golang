package service

import (
	"fmt"

	timeHelper "github.com/kelompok43/Golang/helpers/time"
	"github.com/kelompok43/Golang/payment_method/domain"
)

type paymentMethodService struct {
	repository domain.Repository
}

// DeleteData implements domain.Service
func (pms paymentMethodService) DeleteData(id int) (err error) {
	errResp := pms.repository.Delete(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// UpdateData implements domain.Service
func (pms paymentMethodService) UpdateData(id int, domain domain.PaymentMethod) (paymentMethodObj domain.PaymentMethod, err error) {
	payment_method, errGetByID := pms.GetByID(id)

	if errGetByID != nil {
		return payment_method, errGetByID
	}

	fmt.Println(payment_method)

	domain.CreatedAt = payment_method.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	paymentMethodObj, err = pms.repository.Update(id, domain)

	if err != nil {
		return paymentMethodObj, err
	}

	return paymentMethodObj, nil
}

// GetByID implements domain.Service
func (pms paymentMethodService) GetByID(id int) (paymentMethodObj domain.PaymentMethod, err error) {
	paymentMethodObj, err = pms.repository.GetByID(id)

	if err != nil {
		return paymentMethodObj, err
	}

	return paymentMethodObj, nil
}

func (pms paymentMethodService) InsertData(domain domain.PaymentMethod) (paymentMethodObj domain.PaymentMethod, err error) {
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	paymentMethodObj, err = pms.repository.Create(domain)

	if err != nil {
		return paymentMethodObj, err
	}

	return paymentMethodObj, nil
}

// GetAllData implements domain.Service
func (pms paymentMethodService) GetAllData() (paymentMethodObj []domain.PaymentMethod, err error) {
	paymentMethodObj, _ = pms.repository.Get()

	if err != nil {
		return paymentMethodObj, err
	}

	return paymentMethodObj, nil
}

func NewPaymentMethodService(repo domain.Repository) domain.Service {
	return paymentMethodService{
		repository: repo,
	}
}
