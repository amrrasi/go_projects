package main

import "fmt"

func CalculatePriceRooms(RoomType string, Nights int, PersonCount int) (FinalPrice int) {
	price := 0
	switch RoomType {
	case "Standard":
		price = 1000 * Nights * PersonCount
	case "Suite":
		price = 2000 * Nights * PersonCount
	case "Double":
		price = 3000 * Nights * PersonCount
	default:
		fmt.Printf("Room Not Found")

	}
	tax := float64(price) * 0.09
	FinalPrice = int(tax) + price

	return FinalPrice
}
