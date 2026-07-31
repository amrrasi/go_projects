package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	response, err := http.Get("https://dummyjson.com/products/categories")
	if err != nil {
		fmt.Println("an error has occured ")

	}
	println(response.Status)
	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		println("an error has occured on reading response body")
		fmt.Println(string(responseBody))
	}
	defer response.Body.Close()
}
