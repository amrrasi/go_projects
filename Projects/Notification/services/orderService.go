package services

import (
	"Notification/entities"
	"Notification/externalServices"
	"fmt"
)

type OrderService struct {
	emailService *externalServices.EmailService
	smsService   *externalServices.SmsService
}

func (o *OrderService) CreateOrder(order *entities.Order) *entities.Order {
	fmt.Printf("Created order! %v\n", order)
	o.emailService.SendMessage(order)
	o.smsService.SendMessage(order)
	return order
}

func NewOrderService() *OrderService {
	return &OrderService{
		emailService: externalServices.NewEmailService(),
		smsService:   externalServices.NewSmsService(),
	}
}
