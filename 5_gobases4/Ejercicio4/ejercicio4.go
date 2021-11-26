package main

import (
	"errors"
	"fmt"
)

func calcSalary(h int, s float64) (float64, error) {
	total := float64(h) * s
	if h < 80 {
		err := errors.New("error: el trabajador no puede haber trabajado menos de 80h mesuales")
		return 0.0, err
	}
	if total >= 150.000 {
		total = total - (total * 0.1)
		return total, nil
	}
	return total, nil
}
func main() {
	fmt.Println()
	salary, err := calcSalary(81, 160.000)
	if err != nil { // si err no es nulos es por que existe un error
		fmt.Println(err)
	} else {
		fmt.Printf("El salario a pagar es: %.3f", salary)
	}
	fmt.Println()
}
