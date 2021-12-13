package main

import (
	"errors"
	"fmt"
)

const (
	a_cat = "A"
	b_cat = "B"
	c_cat = "C"
)

func main() {
	your_salary, err := salary_calculator(10800.00, a_cat)
	if err == nil {
		fmt.Printf("Su salario es de: %.2f\n", your_salary)
	} else {
		fmt.Printf("Error: %v", err)
	}
}

func salary_calculator(minutes_worked float64, category string) (float64, error) {
	switch {
	case category == a_cat:
		return salary_coordinator(minutes_worked, a_cat_salary), nil
	case category == b_cat:
		return salary_coordinator(minutes_worked, b_cat_salary), nil
	case category == c_cat:
		return salary_coordinator(minutes_worked, c_cat_salary), nil
	default:
	}
	return 0, errors.New("no values were introduced")
}

func salary_coordinator(minutes_worked float64, operation func(time_worked float64) float64) float64 {
	return operation(minutes_worked)
}

func a_cat_salary(minutes_worked float64) float64 {
	salary := (minutes_worked / 60) * 3000
	salary = salary + (salary * 0.5)
	return salary
}

func b_cat_salary(minutes_worked float64) float64 {
	salary := (minutes_worked / 60) * 1500
	salary = salary + (salary * 0.2)
	return salary
}

func c_cat_salary(minutes_worked float64) float64 {
	return (minutes_worked / 60) * 1000
}
