package domain

type Repository interface {
	SendEmail(toAddress, title, message, media string) (Hourmailer, error)
}
