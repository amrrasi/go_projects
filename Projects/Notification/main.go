package main

import (
	"Notification/entities"
	"Notification/services"
	"fmt"
)

func main() {

	orderAmir := entities.Order{
		ID:           1,
		UserFullName: "amirreza",
		UserPhone:    "09930007410",
		Price:        float64(100),
		Status:       true,
	}

	orderService := services.NewOrderService()
	orderService.CreateOrder(&orderAmir)
	fmt.Printf("Order Created! %v\n", orderAmir)

}
