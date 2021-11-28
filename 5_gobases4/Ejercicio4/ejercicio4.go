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
func halftBonus(s float64, m int) (float64, error) {
	total := float64(m) * (s / 12)
	if m <= 0 || s <= 0 {
		err := errors.New("error: los datos no pueden ser negativos")
		return 0.0, err
	}
	return total, nil
}
func main() {
	fmt.Println()
	// punto a
	salary, err := calcSalary(81, 160.000)
	if err != nil { // si err no es nulos es por que existe un error
		fmt.Println(err)
	} else {
		fmt.Printf("El salario a pagar es: %.3f", salary)
	}
	// fin punto a
	fmt.Println()
	// punto b
	bonus, e := halftBonus(salary, -1)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("El bono a pagar es: %.3f", bonus)
	}
	// fin punto b
	fmt.Println()
}
