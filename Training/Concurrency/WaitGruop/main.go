package main

import (
	"github.com/go-test/deep"
	"io"
	"net/http"
)

var Todolist = []string{}

func main() {

}

func getUrl() {
	response, err := http.Get("https://fanapit.com")
	if err != nil {
		panic(error)
	}

	responseBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		panic(error)
	}

	Todolist = append(Todolist, string(responseBody))
}
