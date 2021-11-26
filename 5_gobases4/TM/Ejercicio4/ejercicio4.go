package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	salary, err := calculateSalary(80, 150000)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("El salario es: %.2f\n", salary)

	bonus, err := calculateHalfBonus(salary, 1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("El aguinaldo es: %.2f\n", bonus)
}

type SalaryError struct {
	Status  int
	Message string
}

func (e *SalaryError) Error() string {
	return fmt.Sprintf("status: %d, message: %s", e.Status, e.Message)
}

func calculateSalary(hoursWorked int, pricePerHour float64) (float64, error) {
	if hoursWorked < 80 {
		return 0.0, errors.New("error: el trabajador no puede ser haber trabajado menos de 80 horas")
	}
	if pricePerHour <= 0 {
		return 0.0, fmt.Errorf("error: el salario por hora no puede ser negativo o 0. Precio ingresado: %.2f", pricePerHour)
	}

	salary := float64(hoursWorked) * pricePerHour

	if salary >= 150000 {
		salary = salary * 0.9
	}

	return salary, nil
}

func calculateHalfBonus(bestSalaryOfSemester float64, monthsWorked int) (float64, error) {
	if bestSalaryOfSemester <= 0 {
		err := fmt.Errorf("se genero un error, %w ", &SalaryError{
			Status:  500,
			Message: "El salario no puede ser menor o igual a 0",
		})

		return 0.0, errors.Unwrap(err)
	}
	if monthsWorked <= 0 {
		err := fmt.Errorf("se genero un error, %w ", &SalaryError{
			Status:  500,
			Message: "El trabajador debe haber trabajado al menos 1 mes",
		})

		return 0.0, errors.Unwrap(err)
	}

	return bestSalaryOfSemester / 12 * float64(monthsWorked), nil
}
