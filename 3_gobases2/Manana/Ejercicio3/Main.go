package main

import (
	"errors"
	"fmt"
)

func calcularSalario(minutos int, categoria string) (float64, error) {

	horasTrabajadas := float64(minutos) / 60
	fmt.Println("Las horas trabajadas son", horasTrabajadas)
	var salarioFinal float64
	switch categoria {
	case "C":
		salarioFinal = float64(horasTrabajadas) * 1000.0
	case "B":
		total := float64(horasTrabajadas) * 1500
		salarioFinal = total + total*0.2
	case "A":
		total := float64(horasTrabajadas) * 3000.0
		salarioFinal = total + total*0.5
	default:
		return 0.0, errors.New("categoria no valida")
	}

	return salarioFinal, nil
}

func main() {

	fmt.Println("Bienvenidos al ejercicio 3")

	fmt.Println("CASO MARINERO1")
	resultado1, err := calcularSalario(30, "B")
	if err == nil {
		fmt.Println(resultado1)
	} else {
		fmt.Println(err)
	}

	fmt.Println("CASO MARINERO2")
	resultado2, err := calcularSalario(75, "T")
	if err == nil {
		fmt.Println(resultado2)
	} else {
		fmt.Println(err)
	}

	fmt.Println("CASO MARINERO3")
	resultado3, err := calcularSalario(30, "C")
	if err == nil {
		fmt.Println(resultado3)
	} else {
		fmt.Println(err)
	}
}
