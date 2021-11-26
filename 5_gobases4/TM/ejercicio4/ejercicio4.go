package main

import (
	"errors"
	"fmt"
)

func SalaryHoursCalculator(hours int, value float64) (float64, error) {
	if hours > 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	var salary float64 = float64(hours) * value
	if salary >= 150000 {
		salary -= salary * 0.1
	}
	return salary, nil
}

func Bonus(bestSalary float64, workedMonths int) (float64, error) {
	if bestSalary <= 0 || workedMonths <= 0 {
		return 0, fmt.Errorf("invalid parameters: %v, %v", bestSalary, workedMonths)
	}
	bonus := bestSalary / 12 * float64(workedMonths)
	return bonus, nil
}

func Ej4(option string, hours, workedMonths int, value, bestSalary float64) (float64, error) {

	switch option {
	case "salary":
		salary, err := SalaryHoursCalculator(hours, value)
		if err != nil {
			return 0.0, err //errors.Unwrap(err)
		}
		return salary, nil
	case "bonus":
		bonus, err := Bonus(bestSalary, workedMonths)
		if err != nil {
			return 0.0, err //errors.Unwrap(err)
		}
		return bonus, nil
	default:
		return 0.0, fmt.Errorf("invalid option %s", option)
	}
}

func main() {
	fmt.Println(Ej4("salary", 80, 0, 10.0, 0.0))
	fmt.Println(Ej4("bonus", 0.0, -6, 0.0, 150000))
	fmt.Println(Ej4("mal", 80, 0, 10.0, 0.0))
	// salary, err := SalaryHoursCalculator(80, 10.0)
	// bonus, err := Bonus(150000, -7)
	// if err != nil {
	// 	e := fmt.Errorf("Hubo un error")
	// 	fmt.Println(e)
	// 	err = e
	// 	fmt.Println(errors.Unwrap((err)))
	// } else {
	// 	fmt.Println(salary)
	// 	fmt.Println(bonus)
	// }
}
