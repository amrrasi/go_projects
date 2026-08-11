package main

import "encoding/json"

type Person struct {
	Name         string `json:"name"`
	Family       string `json:"family"`
	Age          int    `json:"age,omitempty"`
	NationalCode string `json:"nationalCode"`
}

func main() {

	Person1 := Person{Name: "Amir", Family: "Askari", Age: 22, NationalCode: "0250248786"}
	Person2 := Person{Name: "Ali", Family: "izi", NationalCode: "1111111111"}
	Person3 := Person{Name: "Bardia", Family: "Biggy", Age: 10, NationalCode: "2222222222"}
	Person4 := Person{Name: "Mamad", Family: "Seifi", Age: 66, NationalCode: "3333333333"}

	Persons := []Person{Person1, Person2, Person3, Person4}
	personJson, err := json.Marshal(Persons)
	if err != nil {
		panic(err)
	}

	println(string(personJson))

}
