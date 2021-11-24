package main

import (
	"errors"
	"fmt"
)

func calcSalario(minutos int, categoria string) (float64, error) {
	var sueldoCategorias = map[string]float64{"A": 3000, "B": 1500, "C": 1000}
	var aumentoCategorias = map[string]float64{"A": 0, "B": 1.2, "C": 1.5}

	if sueldoPorHora, ok := sueldoCategorias[categoria]; ok {
		sueldoHoras := (float64(minutos) / 60) * sueldoPorHora
		return sueldoHoras * aumentoCategorias[categoria], nil
	}

	return 0, errors.New("Se ingreso una categoria no aceptada")
}

func pruebaSalario(minutos int, categoria string) {
	salario, err := calcSalario(minutos, categoria)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario es %.2f\n", salario)
	}
}

func main() {
	pruebaSalario(8000, "A")
	pruebaSalario(8000, "B")
	pruebaSalario(8000, "C")
	pruebaSalario(8000, "D")
}
