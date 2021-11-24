package main

import (
	"fmt"
)

func calculate_hours(min_work int) float64 {
	return (float64)(min_work / 60)
}

func calculate_salary(min_worked int, cat string) (float64, error) {
	var salary float64

	switch cat {
	case "C":
		salary = 1000 * calculate_hours(min_worked)
	case "B":
		salary = (1500 * calculate_hours(min_worked)) + (0.20 * (1500 * calculate_hours(min_worked)))
	case "A":
		salary = (3000 * calculate_hours(min_worked)) + (0.50 * (3000 * calculate_hours(min_worked)))
	}

	return salary, nil
}

func main() {
	var min_worked int = 678
	var cat string = "A"
	salary, err := calculate_salary(min_worked, cat)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario correspondiente es de %6.2f", salary)
	}
}
