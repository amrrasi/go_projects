package main

import (
	"fmt"
	"github.com/amrrasi/go_projects/services"
	jalali "github.com/jalaali/go-jalaali"
)

func main() {

	fmt.Printf("Hello World! \n")
	year, month, day, error := jalali.ToGregorian(1404, 06, 05)
	if error == nil {
		fmt.Printf("Shamsi Year is %d/%d/%d\n", year, month, day)
	} else {
		fmt.Printf("Shamsi Year ba moshkel movajeh shode\n")
	}

	var services services.TestService = services.TestService{}
	fmt.Printf("%v", services)

}
