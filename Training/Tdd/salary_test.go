package main

import "testing"

func TestExtraSalaryCalculate(t *testing.T) {
	baseSalary := 20000000
	extraHours := 12
	want := 20000012

	Got := ExtraSalaryCalculate(baseSalary, extraHours)

	if want != Got {
		t.Errorf("got %v instead of %v", want, Got)
	}
}
