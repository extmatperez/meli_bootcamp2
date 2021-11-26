package main

import (
	"fmt"
	"errors"
)

func Ej4(option string, hours, monthsWorking int, valuePerHour, bestSalary float64) (float64, error) {

	switch option {
	case "Salary":
		salary, err := Salary(hours, valuePerHour)
		if err != nil {
			return 0, errors.Unwrap(err)
		}
		return salary, nil

	case "Bonus":
		bonus, err := Bonus(bestSalary, monthsWorking)
		if err != nil {
			e1 := fmt.Errorf("nuevo error: %w", err)
			fmt.Println(e1)
			return 0, errors.Unwrap(e1)
		}
		return bonus, nil

	default:
		return 0.0, fmt.Errorf("invalid option %s", option)
	}

}

func Salary(hours int, valuePerHour float64) (float64, error) {

	if hours < 80 {
		err1 := errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		err2 := fmt.Errorf("error 2: %w", err1)
		return 0.0, err2
	}
	salary := float64(hours) * valuePerHour

	if salary >= 150000 {
		salary -= salary * 0.1
	}

	return salary, nil

}

func Bonus(bestSalary float64, monthsWorked int) (float64, error) {
	
	if bestSalary < 0 || monthsWorked < 0 {
		return 0.0, errors.New("error: invalid parameter")
	}

	bonus := bestSalary / 12 * float64(monthsWorked)

	return bonus, nil
}


func main() {
	fmt.Println(Ej4("Salary", 80, 0, 10.0, 0.0))
	fmt.Println(Ej4("Bonus", 0.0, 6, 0, 150000))
}