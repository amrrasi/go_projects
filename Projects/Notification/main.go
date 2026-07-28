package main

import (
	"Notification/entities"
	"Notification/externalServices"
	"Notification/services"
)

func main() {

	orderAmir := entities.Order{
		ID:           1,
		UserFullName: "amirreza",
		UserId:       "55",
		UserPhone:    "09930007410",
		Price:        float64(100),
		Status:       true,
	}

	orderService := services.NewOrderService(externalServices.NewSmsService())
	orderService.CreateOrder(&orderAmir)
}
