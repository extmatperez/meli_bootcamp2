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
		return fmt.Errorf("salary is less than 150000 (%d)", (salary))
	}
	return nil
}