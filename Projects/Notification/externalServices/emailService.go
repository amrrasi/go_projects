package externalServices

import (
	"Notification/entities"
	"fmt"
)

type EmailService struct{}

func (e *EmailService) SendMessage(order *entities.Order) {
	fmt.Printf("Email has been Sent! %v\n", order)
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
