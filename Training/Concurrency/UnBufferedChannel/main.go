package main

func main() {
	numChannel := make(chan int)

	go func() {
		numChannel <- 1
	}()

	receivedChannel := <-numChannel
	println(receivedChannel)
}
