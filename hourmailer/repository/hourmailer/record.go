package repoHourmailer

import (
	"github.com/kelompok43/Golang/hourmailer/domain"
)

type Response struct {
	Message string `json:"message"`
}

func (res *Response) toDomain() domain.Hourmailer {
	return domain.Hourmailer{
		Message: res.Message,
	}
}
