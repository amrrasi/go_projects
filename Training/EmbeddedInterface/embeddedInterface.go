package main

import "fmt"

type Animal interface {
	Sleep()
	Eat()
	Walk()
}

type Human interface {
	Animal
	Speak()
	Think()
}

type Athlete struct {
	name string
	age  int
}

type Dog struct {
	name string
}

func main() {

	athlete1 := &Athlete{name: "Amirreza", age: 22}
	dog1 := &Dog{name: "dog"}

	var human Human = athlete1
	var animal Animal = dog1

	human.Eat()
	human.Walk()
	human.Speak()
	human.Think()
	human.Sleep()

	animal.Walk()
	animal.Eat()
	animal.Sleep()

}

func (dog *Dog) Eat() {
	fmt.Printf("Dog is Eating...")
}
func (dog *Dog) Sleep() {
	fmt.Printf("Dog is Sleeping...")
}
func (dog *Dog) Walk() {
	fmt.Printf("Dog is Walking...")
}

func (athlete *Athlete) Eat() {
	fmt.Printf("Athlete is Eating...")
}
func (athlete *Athlete) Sleep() {
	fmt.Printf("Athlete is Sleeping...")
}
func (athlete *Athlete) Walk() {
	fmt.Printf("Athlete is Walking...")
}
func (athlete *Athlete) Think() {
	fmt.Printf("Athlete is Thinking...")
}
func (athlete *Athlete) Speak() {
	fmt.Printf("Athlete is Speaking...")
}
