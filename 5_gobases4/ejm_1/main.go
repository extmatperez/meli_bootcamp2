package main

import (
	"fmt"
)

func main() {
	salary := 140000
	err := checkSalary(salary)
	fmt.Println("error:", err)
}

func checkSalary(salary int) error {
	if salary < 150000 {
		return SalaryError{salary}
	}
	return nil
}

type SalaryError struct {
	Salary int
}

func (e SalaryError) Error() string {
	return fmt.Sprintf("salary is less than 150000 (%d)", (e.Salary))
}