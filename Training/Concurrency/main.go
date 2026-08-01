package main

import "time"

func main() {

	value := 0

	go Task1()
	go Task2()
	go Task3()

	go func() {
		value++
	}()

	println(value)
	time.Sleep(time.Second)

}

func Task1() {
	println("Task1")
}

func Task2() {
	println("Task2")
}

func Task3() {
	println("Task3")
}
