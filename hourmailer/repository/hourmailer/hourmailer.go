package repoHourmailer

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	url := "https://hourmailer.p.rapidapi.com/send"

	payload := strings.NewReader(fmt.Sprintf("{\r\"toAddress\": \"%s\",\r\"title\": \"%s\",\r\"message\": \"<p><img alt=picture width=200 src=%s /></p> %s\"\r}", toAddress, title, media, message))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "531f7adc42msha9bdbd410bfc903p13cc07jsnf1ffc381bb61")
	req.Header.Add("X-RapidAPI-Host", "hourmailer.p.rapidapi.com")

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
