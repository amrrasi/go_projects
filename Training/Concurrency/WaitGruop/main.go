package main

import (
	//"github.com/go-test/deep"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var Todolist = []string{}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		GetTodo(i, &wg)
	}
	fmt.Printf("%v", Todolist)

	wg.Wait()
}

func GetTodo(id int, wg *sync.WaitGroup) {
	GetUrl("https://jsonplaceholder.typicode.com/todos/"+strconv.Itoa(id), wg)
}

func GetUrl(url string, wg *sync.WaitGroup) {

	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	responseBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	Todolist = append(Todolist, string(responseBody))
}
