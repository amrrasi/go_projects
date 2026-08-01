package main

import (
	"log"
	"os"
)

var (
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
)

func init() {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Failed to open long file", err)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetFlags(flags)
	log.SetOutput(file)
	ErrorLogger = log.New(file, "error: ", flags)
	InfoLogger = log.New(file, "Info: ", flags)
	WarnLogger = log.New(file, "Warning: ", flags)

}

func main() {

	ErrorLogger.Println("Start of main")
	Sum(1, 2)
	ErrorLogger.Println("End of main")

}

func Sum(a, b int) {
	InfoLogger.Println("Start ovf Sum")
	println(a + b)
	InfoLogger.Println("End of Sum")
}
