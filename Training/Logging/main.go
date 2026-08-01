package main

import (
	"log"
	"os"
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Failed to open long file", err)
	}
	log.SetOutput(file)
}

func main() {

	log.Println("Start of main")
	Sum(1, 2)
	log.Println("End of main")

}

func Sum(a, b int) {
	log.Println("Start ovf Sum")
	println(a + b)
	log.Println("End of Sum")
}
