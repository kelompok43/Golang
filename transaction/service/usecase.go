package service

import (
	"bytes"
	"fmt"
	"io"

	storageHelper "github.com/kelompok43/Golang/helpers/azure"
	timeHelper "github.com/kelompok43/Golang/helpers/time"
	membershipDomain "github.com/kelompok43/Golang/membership/domain"
	"github.com/kelompok43/Golang/transaction/domain"
	userDomain "github.com/kelompok43/Golang/user/domain"
)

type transactionService struct {
	repository        domain.Repository
	membershipService membershipDomain.Service
	userService       userDomain.Service
}

// DeleteData implements domain.Service
func (ts transactionService) DeleteData(id int) (err error) {
	errResp := ts.repository.Delete(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// UpdateData implements domain.Service
func (ts transactionService) UpdateStatus(id int, domain domain.Transaction) (transactionObj domain.Transaction, err error) {
	transaction, errGetByID := ts.GetByID(id)

	if errGetByID != nil {
		return transaction, errGetByID
	}

	fmt.Println(transaction)

	//add user to become a membership
	if domain.Status == "Diterima" {
		_, err := ts.membershipService.InsertOrder(transaction.ID, transaction.TotalPrice)

		if err != nil {
			return transactionObj, err
		}

		//change status user
		_, err = ts.userService.UpdateStatus(transaction.UserID)

		if err != nil {
			return transactionObj, err
		}
	}

	transaction.Status = domain.Status
	transaction.UpdatedAt = timeHelper.Timestamp()
	transactionObj, err = ts.repository.Update(id, transaction)

	if err != nil {
		return transactionObj, err
	}

	return transactionObj, nil
}

// GetByID implements domain.Service
func (ts transactionService) GetByID(id int) (transactionObj domain.Transaction, err error) {
	transactionObj, err = ts.repository.GetByID(id)

	if err != nil {
		return transactionObj, err
	}

	return transactionObj, nil
}

func (ts transactionService) InsertData(domain domain.Transaction) (transactionObj domain.Transaction, err error) {
	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, domain.Payment_Receipt); err != nil {
		return transactionObj, err
	}

	data := buf.Bytes()
	domain.Status = "Sedang Diproses"
	domain.PictureLink, _ = storageHelper.UploadBytesToBlob(data)
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	transactionObj, err = ts.repository.Create(domain)

	if err != nil {
		return transactionObj, err
	}

	return transactionObj, nil
}

// GetAllData implements domain.Service
func (ts transactionService) GetAllData() (transactionObj []domain.Transaction, err error) {
	transactionObj, _ = ts.repository.Get()

	if err != nil {
		return transactionObj, err
	}

	return transactionObj, nil
}

func NewTransactionService(repo domain.Repository, ms membershipDomain.Service, us userDomain.Service) domain.Service {
	return transactionService{
		repository:        repo,
		membershipService: ms,
		userService:       us,
	}
}
