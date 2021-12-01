package main

import (
	"fmt"
)

type myError struct {
	salary int
	msg    string
}

func (e *myError) Error() string {
	return fmt.Sprintf("for %d - %v", e.salary, e.msg)
}

func taxErrors(salary int) (string, error) {
	if salary < 150000 {
		return salary, myError{
			salary: salary,
			msg:    "El salario no alcanza el minimo imponible",
		}
	}
}

func main() {

}

// ejercicio 4
// func Ej4(option string, hours, monthsWorking int, valuePerHour, bestSalary float64) (float64, error) {
// 	switch option {
// 	case "Salary":
// 		salary, err := Salary(hours, valuePerHour)
// 		if err != nil {
// 			return 0.0, errors.Unwrap(err)
// 		}

// 		return salary, nil
// 	case "Bonus":
// 		bonus, err := Bonus(bestSalary, monthsWorking)
// 		if err != nil {
// 			return 0.0, errors.Unwrap(err)
// 		}

// 		return bonus, nil
// 	default:
// 		return 0.0, fmt.Errorf("invalid option %s", option)
// 	}
// }
// func Salary(hours int, valuePerHour float64) (float64, error) {
// 	if hours < 80 {
// 		return 0.0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
// 	}

// 	salary := float64(hours) * valuePerHour

// 	if salary >= 150000 {
// 		salary -= salary * 0.1
// 	}

// 	return salary, nil
// }
// func Bonus(bestSalary float64, monthsWorking int) (float64, error) {
// 	if bestSalary < 0 || monthsWorking < 0 {
// 		return 0.0, errors.New("invalid parameter")
// 	}

// 	bonus := bestSalary / 12 * float64(monthsWorking)

// 	return bonus, nil
// }
