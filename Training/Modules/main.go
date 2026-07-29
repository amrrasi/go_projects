package main

import (
	"fmt"
	jalali "github.com/jalaali/go-jalaali"
)

func main() {

	fmt.Printf("Hello World! \n")
	year, month, day, error := jalali.ToGregorian(1404, 06, 05)
	if error == nil {
		fmt.Printf("Shamsi Year is %d/%d/%d", year, month, day)
	} else {
		fmt.Printf("Shamsi Year ba moshkel movajeh shode")
	}

}
