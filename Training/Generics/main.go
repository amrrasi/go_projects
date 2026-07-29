package main

import "fmt"

type Number interface {
	int | int64 | float64 | float32
}

func main() {

	x := Sum(2.2, 7.7)
	fmt.Printf("%f\n", x)

	mySlice := []int{1, 2, 3, 4, 5, 6, 7}
	mySlice2 := []float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1}
	fmt.Printf("%d\n", SliceSum(mySlice))
	fmt.Printf("%f\n", SliceSum(mySlice2))
}

func Sum[T Number](a, b T) T {
	return a + b

}

func SliceSum[T Number](slc []T) (sum T) {
	for _, v := range slc {
		sum += v
	}
	return sum
}
