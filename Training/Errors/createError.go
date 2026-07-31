package main

import (
	"errors"
	"fmt"
	//"io"
	//"net/http"
)

func main() {

	output, err := DisplayErrors(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)

}

func DisplayErrors(Number int) (int, error) {

	if Number == 0 {
		return 0, errors.New("number is not valid it is 0")
	}
	return Number * 2, nil

}

func DisplayErrors2(Number int) (int, error) {

	if Number == 0 {
		return 0, fmt.Errorf("number is not valid it is 0 %d", Number)
	}
	return Number * 2, nil

}
