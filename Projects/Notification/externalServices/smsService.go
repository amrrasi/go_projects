package externalServices

import (
	"Notification/entities"
	"fmt"
)

type SmsService struct{}

func (e *SmsService) SendMessage(order *entities.Order) {
	fmt.Printf("SMS has been Sent! %v\n", order)
}

func NewSmsService() *SmsService {
	return &SmsService{}
}
