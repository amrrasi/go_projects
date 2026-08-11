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

	persons := []Person{Person1, Person2}
	PersonJson, err := json.Marshal(persons)
	if err != nil {
		panic(err)
	}

	println(string(PersonJson))

}
