package main

import (
	"fmt"
	"genericList/generics"
)

func main() {
	GenericInt()
	GenericString()
}

func GenericInt() {
	list1 := generics.List[int]{Items: []int{}}
	list1.Add(11)
	list1.Add(22)
	list1.Add(33)
	list1.Add(44)
	list1.Add(55)
	list1.Add(66)
	fmt.Printf("%v\n", list1.Items)

	list1.InsertAt(196, 3)
	fmt.Printf("%v\n", list1.Items)

	list1.RemoveItem(0)
	fmt.Printf("%v\n", list1.Items)

	list1.Remove(55)
	fmt.Printf("%v\n", list1.Items)
}

func GenericString() {
	list1 := generics.List[string]{Items: []string{}}
	list1.Add("Ali")
	list1.Add("Amir")
	list1.Add("Mamad")
	list1.Add("Reza")
	list1.Add("Bardia")
	list1.Add("Sadra")
	fmt.Printf("%v\n", list1.Items)

	list1.InsertAt("Sadegh", 3)
	fmt.Printf("%v\n", list1.Items)

	list1.RemoveItem(0)
	fmt.Printf("%v\n", list1.Items)

	list1.Remove("Sadra")gi
	fmt.Printf("%v\n", list1.Items)
}
