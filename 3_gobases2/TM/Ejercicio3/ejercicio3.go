package main

import (
	"errors"
	"fmt"
)

func main() {
	minutesWorked := 120
	category := 'A'
	salary, err := calculateSalary(minutesWorked, category)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario del empleado es: %.2f\n", salary)
	}
}

func calculateSalary(minutesWorked int, category rune) (float64, error) {
	var hours float64 = float64(minutesWorked) / 60

	if hours <= 0 {
		return 0.0, errors.New("el empleado no registra horas trabajadas")
	}

	var salaryPerHour float64 = 0
	var additionalPercentage float64 = 1.00

	switch category {
	case 'C':
		salaryPerHour = 1000
	case 'B':
		salaryPerHour = 1500
		additionalPercentage = 1.20
	case 'A':
		salaryPerHour = 3000
		additionalPercentage = 1.50
	default:
		return 0.0, errors.New("la categoria no existe")
	}

	return hours * salaryPerHour * additionalPercentage, nil
}
