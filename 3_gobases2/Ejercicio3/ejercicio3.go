package main

import (
	"errors"
	"fmt"
)

const (
	SalC       = 1000.0
	SalB       = 1500.0
	SalA       = 3000.0
	hMensuales = 160
)

func salario(cat string, min float64) (float64, error) {
	var salario float64
	var horas float64
	if min < 0 {
		return 0, errors.New("el tiempo no puede ser negativo")
	} else {
		horas = min / 60
	}
	switch cat {
	case "C", "c":
		salario = SalC * horas
	case "B", "b":
		salario = (SalB * horas) + (SalB * hMensuales * 0.2)
	case "A", "a":
		salario = (SalA * horas) + (SalA * hMensuales * 0.5)
	default:
		return 0, errors.New("la categoria no existe")
	}

	return salario, nil
}

func main() {

	var catIn string
	var minIn float64
	fmt.Print("Cual la categoria: ")
	fmt.Scanf("%s", &catIn)
	fmt.Print("Cuanto tiempo trabajo en minutos: ")
	fmt.Scanf("%f", &minIn)

	sal, err := salario(catIn, minIn)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("El Salario es: ", sal)
	}
}
