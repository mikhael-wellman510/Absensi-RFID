package usecases

import (
	"errors"
	"log"
	"time"
)

type (
	EmailService interface {
		Email(name string) string
	}

	emailService struct {
	}
)

func NewEmailService() EmailService {
	return &emailService{}
}

func (e emailService) Email(name string) string {

	go func(email string) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		if _, err := sendEmail(name); err != nil {
			log.Println(err.Error())
		}
	}(name)

	return "Email service Terkirim"

}

func sendEmail(name string) (string, error) {

	if name == "" {
		return "", errors.New("name is required")
	}

	time.Sleep(5 * time.Second)
	log.Printf("Email telah terkirim : %s", name)
	return "Successfully sent email", nil
}
