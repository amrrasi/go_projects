package main

import (
	"fmt"
	"time"
)

func main() {
	numChannel := make(chan int)

	go SendDataWithChannel(numChannel)

	receivedNum := <-numChannel
	PrintlnWithTime("Received number: ", receivedNum)
	receivedNum = <-numChannel
	PrintlnWithTime("Received number: ", receivedNum)
	receivedNum = <-numChannel
	PrintlnWithTime("Received number: ", receivedNum)
	time.Sleep(time.Second * 2)
}

func SendDataWithChannel(numChannel chan int) {
	PrintlnWithTime("before 1 ")
	numChannel <- 1
	PrintlnWithTime("After 1 ")

	PrintlnWithTime("before 2 ")
	numChannel <- 2
	PrintlnWithTime("After 2 ")

	PrintlnWithTime("before 3 ")
	numChannel <- 3
	PrintlnWithTime("After 3 ")
}

func PrintlnWithTime(args ...any) {
	fmt.Printf("Time: %s, %v\n", time.Now().Format(time.RFC3339Nano), args)
}
