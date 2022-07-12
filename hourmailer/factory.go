package hourmailer

import (
	domainHourmailer "github.com/kelompok43/Golang/hourmailer/domain"
	repoHourmailer "github.com/kelompok43/Golang/hourmailer/repository/hourmailer"
)

func NewHourmailerFactory() domainHourmailer.Repository {
	return repoHourmailer.NewHourmailerRepository()
}
