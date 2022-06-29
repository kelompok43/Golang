package service

import (
	"bytes"
	"fmt"
	"io"

	storageHelper "github.com/kelompok43/Golang/helpers/azure"
	timeHelper "github.com/kelompok43/Golang/helpers/time"
	"github.com/kelompok43/Golang/transaction/domain"
)

type transactionService struct {
	repository domain.Repository
}

// UpdateStatus implements domain.Service
func (transactionService) UpdateStatus(id int, status string) (transactionObj domain.Transaction, err error) {
	panic("unimplemented")
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
func (ts transactionService) UpdateData(id int, domain domain.Transaction) (transactionObj domain.Transaction, err error) {
	transaction, errGetByID := ts.GetByID(id)

	if errGetByID != nil {
		return transaction, errGetByID
	}

	fmt.Println(transaction)

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

func NewTransactionService(repo domain.Repository) domain.Service {
	return transactionService{
		repository: repo,
	}
}
