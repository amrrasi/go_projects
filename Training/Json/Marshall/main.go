package main

import "encoding/json"

type Person struct {
	Name   string `json:"name"`
	Family string `json:"family"`
	Age    int    `json:"age,omitempty"`
}

func main() {

	Person1 := Person{Name: "Amir", Family: "Askari", Age: 22}
	Person2 := Person{Name: "Ali", Family: "izi", Age: 0}

	Person1Json, err := json.Marshal(Person1)
	if err != nil {
		panic(err)
	}

	Person2Json, err := json.Marshal(Person2)
	if err != nil {
		panic(err)
	}

	println(string(Person1Json))
	println(string(Person2Json))

}
