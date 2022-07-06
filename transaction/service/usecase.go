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

func (ts transactionService) GetUserTrx(userID int) (transactionObj []domain.Transaction, err error) {
	transactionObj, err = ts.repository.GetByUserID(userID)

	if err != nil {
		return transactionObj, err
	}

	return transactionObj, nil
}

func (ts transactionService) GetUserTrxByID(userID, trxID int) (transactionObj domain.Transaction, err error) {
	userTransaction, err := ts.repository.GetByUserID(userID)

	if err != nil {
		return transactionObj, err
	}

	for _, value := range userTransaction {
		if value.ID == trxID {
			transactionObj, err = ts.repository.GetByID(trxID)

			if err != nil {
				return transactionObj, err
			}
		}
	}

	return transactionObj, nil
}

// InsertDetail implements domain.Service
func (ts transactionService) InsertDetail(id int, price int) (transactionDetailObj domain.TransactionDetail, err error) {
	var transactionDetail domain.TransactionDetail
	membershipCategory, err := ts.membershipService.GetCategoryByPrice(price)

	if err != nil {
		return transactionDetailObj, err
	}

	transactionDetail.TransactionID = id
	transactionDetail.MembershipCategoryID = membershipCategory.ID
	transactionDetailObj, err = ts.repository.CreateDetail(transactionDetail)

	if err != nil {
		return transactionDetailObj, err
	}

	return transactionDetailObj, nil
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
		trxDetail, err := ts.InsertDetail(transaction.ID, transaction.TotalPrice)

		if err != nil {
			return transactionObj, err
		}

		//change status user
		_, err = ts.membershipService.InsertData(transaction.UserID, trxDetail.MembershipCategoryID)

		if err != nil {
			return transactionObj, err
		}

		fmt.Println("trx userID = ", transaction.UserID)
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
