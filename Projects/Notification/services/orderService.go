package services

import (
	"Notification/entities"
	"Notification/externalServices"
	"fmt"
)

type OrderService struct {
	Notifier externalServices.Notifier
}

func (orderService *OrderService) CreateOrder(order *entities.Order) *entities.Order {
	fmt.Printf("Created order! %v\n", order)
	orderService.Notifier.SendNotify(order.UserId, "order created")

	return order
}

func NewOrderService(notifier externalServices.Notifier) *OrderService {
	return &OrderService{
		Notifier: notifier,
	}
}
