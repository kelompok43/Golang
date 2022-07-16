package service

import (
	"bytes"
	"fmt"
	"io"

	storageHelper "github.com/kelompok43/Golang/helpers/azure"
	timeHelper "github.com/kelompok43/Golang/helpers/time"
	hourmailerDomain "github.com/kelompok43/Golang/hourmailer/domain"
	"github.com/kelompok43/Golang/news/domain"
	userDomain "github.com/kelompok43/Golang/user/domain"
)

type newsService struct {
	repository     domain.Repository
	hourmailerRepo hourmailerDomain.Repository
	userService    userDomain.Service
}

// DeleteCategory implements domain.Service
func (ns newsService) DeleteCategory(id int) (err error) {
	errResp := ns.repository.DeleteCategory(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// GetAllCategory implements domain.Service
func (ns newsService) GetAllCategory() (newsCategoryObj []domain.Category, err error) {
	newsCategoryObj, _ = ns.repository.GetCategory()

	if err != nil {
		return newsCategoryObj, err
	}

	return newsCategoryObj, nil
}

// GetCategoryByID implements domain.Service
func (ns newsService) GetCategoryByID(id int) (newsCategoryObj domain.Category, err error) {
	newsCategoryObj, err = ns.repository.GetCategoryByID(id)

	if err != nil {
		return newsCategoryObj, err
	}

	return newsCategoryObj, nil
}

// InsertCategory implements domain.Service
func (ns newsService) InsertCategory(domain domain.Category) (newsCategoryObj domain.Category, err error) {
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	newsCategoryObj, err = ns.repository.CreateCategory(domain)

	if err != nil {
		return newsCategoryObj, err
	}

	return newsCategoryObj, nil
}

// UpdateCategory implements domain.Service
func (ns newsService) UpdateCategory(id int, domain domain.Category) (newsCategoryObj domain.Category, err error) {
	category, errGetByID := ns.repository.GetCategoryByID(id)

	if errGetByID != nil {
		return newsCategoryObj, errGetByID
	}

	domain.CreatedAt = category.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	newsCategoryObj, err = ns.repository.UpdateCategory(id, domain)

	if err != nil {
		return newsCategoryObj, err
	}

	return newsCategoryObj, nil
}

// DeleteData implements domain.Service
func (ns newsService) DeleteData(id int) (err error) {
	errResp := ns.repository.Delete(id)

	if errResp != nil {
		return errResp
	}

	return nil
}

// UpdateData implements domain.Service
func (ns newsService) UpdateData(id int, domain domain.News) (newsObj domain.News, err error) {
	news, errGetByID := ns.GetByID(id)

	if errGetByID != nil {
		return news, errGetByID
	}

	domain.CreatedAt = news.CreatedAt
	domain.UpdatedAt = timeHelper.Timestamp()
	if domain.Picture != nil {
		buf := bytes.NewBuffer(nil)

		if _, err := io.Copy(buf, domain.Picture); err != nil {
			return newsObj, err
		}

		data := buf.Bytes()
		domain.PictureLink, _ = storageHelper.UploadBytesToBlob(data)
	} else {
		domain.PictureLink = news.PictureLink
	}

	newsObj, err = ns.repository.Update(id, domain)

	if err != nil {
		return newsObj, err
	}

	return newsObj, nil
}

// GetByID implements domain.Service
func (ns newsService) GetByID(id int) (newsObj domain.News, err error) {
	newsObj, err = ns.repository.GetByID(id)

	if err != nil {
		return newsObj, err
	}

	return newsObj, nil
}

func (ns newsService) InsertData(domain domain.News) (newsObj domain.News, err error) {
	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, domain.Picture); err != nil {
		return newsObj, err
	}

	data := buf.Bytes()
	domain.PictureLink, _ = storageHelper.UploadBytesToBlob(data)
	domain.CreatedAt = timeHelper.Timestamp()
	domain.UpdatedAt = timeHelper.Timestamp()
	newsObj, err = ns.repository.Create(domain)

	if err != nil {
		return newsObj, err
	}

	user, err := ns.userService.GetAllData()

	fmt.Println("user = ", user)

	if err != nil {
		return newsObj, err
	}

	userEmail := ""

	for _, value := range user {
		userEmail = value.Email
		fmt.Println("useremail = ", userEmail)
		_, err := ns.hourmailerRepo.SendEmail(userEmail, domain.Title, domain.Description, domain.PictureLink)

		if err != nil {
			return newsObj, err
		}
	}

	return newsObj, nil
}

// GetAllData implements domain.Service
func (ns newsService) GetAllData() (newsObj []domain.News, err error) {
	newsObj, _ = ns.repository.Get()

	if err != nil {
		return newsObj, err
	}

	return newsObj, nil
}

func NewNewsService(repo domain.Repository, hourmailerRepo hourmailerDomain.Repository, userService userDomain.Service) domain.Service {
	return newsService{
		repository:     repo,
		hourmailerRepo: hourmailerRepo,
		userService:    userService,
	}
}
