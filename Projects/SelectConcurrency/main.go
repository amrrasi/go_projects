package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ioutil"
	"net/http"
	"time"
)

func main() {

	GetExample()

}

func GetExample() {
	resource1 := make(chan string)
	resource2 := make(chan string)

	go GetLiveScore(resource1, "https://footba11.co/json/liveFeed")
	go GetLiveScore(resource2, "https://livescore-api.varzesh3.com/v1.0/livescore/today")

	select {
	case result1 := <-resource1:
		println(result1)
	case result2 := <-resource2:
		println(result2)

		PrintlnWithTime("End")

	}
}

func GetLiveScore(content chan<- string, url string) {

	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}
	request.Header = http.Header{}

	request.Header.Add("refer", "https://www.varzesh3.com/")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.body)
	if err != nil {
		panic(err)
	}

	destination := &bytes.Buffer{}

	if err = json.Indent(destination, responseBody, "", "   "); err != nil {
		panic(err)
	}

	PrintlnWithTime("Before set content ...")
	content <- destination.String()
	PrintlnWithTime("After set content ...")

}

func PrintlnWithTime(args ...any) {
	fmt.Printf("Time: %s %v\n", time.Now().Format(time.RFC3339Nano), args)
}
