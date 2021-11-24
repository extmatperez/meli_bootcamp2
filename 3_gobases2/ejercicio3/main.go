package main

import (
	"errors"
	"fmt"
)

func main() {
	salario, err := calcularSalario(60, "A")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario es: %.2f\n", salario)
	}
}

func calcularSalario(min float64, cat string) (float64, error) {
	horas := min / 60
	switch cat {
	case "A":
		return (3000 * horas * 1.5), nil
	case "B":
		return (1500 * horas * 1.2), nil
	case "C":
		return (1000 * horas), nil
	default:
		return 0, errors.New("categoría incorrecta")
	}
}
