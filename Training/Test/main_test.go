package main

import "testing"

func TestCalculatePriceRooms(t *testing.T) {
	//Arrange
	expectedRoomPrice := 4360
	Nights := 2
	PersonCount := 2
	//Act
	actual := CalculatePriceRooms("Standard", Nights, PersonCount)
	//Assert

	if actual != expectedRoomPrice {
		t.Fail()
	}
}
