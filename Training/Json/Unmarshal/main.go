package main

import "encoding/json"

type person struct {
	Name   string
	Family string
	Age    int
}

func main() {
	UnmarshalExample()
}

func UnmarshalExample() {
	person1Json := []byte(`{"name":"Ali","Family":"izi","Age":55}`)
	var person1 = person{}
	err := json.Unmarshal(person1Json, &person1)

	if err != nil {
		panic(err)
	}

	println(person1.Name)
	println(person1.Family)
	println(person1.Age)
}
