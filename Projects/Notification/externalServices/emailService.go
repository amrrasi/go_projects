package externalServices

import (
	"fmt"
)

type EmailService struct{}

func (e *EmailService) SendNotify(receiver string, message string) {
	fmt.Printf("This %s\n Email has been Sent to! %s\n", message, receiver)
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
