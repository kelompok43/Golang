package repoHourmailer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kelompok43/Golang/hourmailer/domain"
)

type hourmailerAPI struct {
	httpClient http.Client
}

var (
	emailTemplate = "template"
)

// SendEmail implements domain.Repository
func (ha hourmailerAPI) SendEmail(toAddress string, title string, message string, media string) (domain.Hourmailer, error) {
	url := os.Getenv("HOURMAILER_URL")

	payload := strings.NewReader(fmt.Sprintf("{\r\"toAddress\": \"%s\",\r\"title\": \"%s\",\r\"message\": \"<h3>%s</h3><p><img alt=picture width=200 src=%s /></p> %s\"\r}", toAddress, title, title, media, message))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", os.Getenv("HOURMAILER_API_KEY"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("HOURMAILER_API_HOST"))

	res, err := ha.httpClient.Do(req)

	if err != nil {
		return domain.Hourmailer{}, err
	}

	defer res.Body.Close()
	data := Response{}
	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		return domain.Hourmailer{}, err
	}

	fmt.Println("response hourmailer = ", data)

	return data.toDomain(), nil
}

func NewHourmailerRepository() domain.Repository {
	return hourmailerAPI{
		httpClient: http.Client{},
	}
}
