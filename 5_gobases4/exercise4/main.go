package main

import (
	"errors"
	"fmt"
)

type Worker struct {
	hours      int
	hourSalary int
	bestSalary int
	workedTime int
	thisError  string
}

func (t *Worker) Error() string {
	return t.thisError
}

func getTotal(w Worker) (float64, error) {

	totalSalary := float64(w.hours * w.hourSalary)

	switch {
	case totalSalary >= 150000:
		return (totalSalary * 0.9), nil
	case w.hourSalary < 80:
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	default:
		return totalSalary, nil
	}
}

func getBonus(w Worker) (float64, error) {
	bonusMount := float64(w.bestSalary) / 12 * float64(w.workedTime)

	if w.bestSalary < 0 || w.workedTime < 0 {
		err := errors.New("error: los valores no pueden ser negativos")
		return 0, fmt.Errorf("yeah: %w", err)
	} else {
		return bonusMount, nil
	}
}

func main() {

	worker1 := Worker{hours: 40, hourSalary: 100, bestSalary: -20201, workedTime: 10}

	response, err := getTotal(worker1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\nEl salario es: ", response)
		response, err := getBonus(worker1)

		if err != nil {
			fmt.Println(errors.Unwrap(err))
		} else {
			fmt.Printf("\nEl bono es: %.2f\n\n", response)
		}
	}
}
