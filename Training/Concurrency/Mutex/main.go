package main

import (
	"fmt"
	"sync"
)

type Employee struct {
	Id     int
	Salary int64
}

var EmployeeSalaryList []Employee

func main() {
	EmployeeAppend()

	remain := PaySalary()

	fmt.Println("Remain:", remain)
}

func getRandomNumber() int64 {
	return 24_000_000
}

func EmployeeAppend() {
	for i := 0; i < 50_000; i++ {
		EmployeeSalaryList = append(EmployeeSalaryList, Employee{
			Id:     i,
			Salary: getRandomNumber(),
		})
	}
}

func PaySalary() int64 {
	var Remain int64 = 150_000_000_000_000

	var wg sync.WaitGroup
	var mx sync.Mutex

	for _, employee := range EmployeeSalaryList {

		wg.Add(1)

		go func(emp Employee) {
			defer wg.Done()

			mx.Lock()
			defer mx.Unlock()

			if Remain >= emp.Salary {
				Remain -= emp.Salary
			}

		}(employee)
	}

	wg.Wait()

	return Remain
}
