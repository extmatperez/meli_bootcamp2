package main

import (
	"errors"
	"fmt"
)

func main() {
	hours := 80
	value := 5000
	salary, err := calcSalary(hours, value)
	fmt.Println("salary:", salary)
	fmt.Println("error:", err)

	bestSalary := float64(-1000)
	months := 12
	prima, err := calcPrima(bestSalary, months)
	fmt.Println("prima:", prima)
	fmt.Println("error:", err)
	bestSalary = salary
	prima, err = calcPrima(bestSalary, months)
	fmt.Println("prima:", prima)
	fmt.Println("error:", err)
}

func calcSalary(hours int, value int) (float64, error) {
	if hours < 80 {
		return 0, errors.New("error: el trabajador no pudo haber trabajado menos de 80 horas al mes")
	}
	var salary float64 = float64(hours * value)
	if salary >= 150000 {
		salary *= 0.9 // 10% taxes discount
	}
	return salary, nil
}

func calcPrima(bestSalary float64, months int) (float64, error) {
	if bestSalary < 0 {
		return 0, NegativeValueError{"bestSalary", bestSalary}
	}
	if months < 0 {
		return 0, NegativeValueError{"months", float64(months)}
	}
	return bestSalary / 12 * float64(months), nil
}

type NegativeValueError struct {
	keyName string
	value   float64
}

func (e NegativeValueError) Error() string {
	return fmt.Sprintf("%s is negative: %.1f", e.keyName, e.value)
}