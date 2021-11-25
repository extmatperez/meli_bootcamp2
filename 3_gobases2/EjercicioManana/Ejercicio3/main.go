package main

import (
	"fmt"
)

const (
	C               float64 = 1000
	B               float64 = 1500
	A               float64 = 3000
	PorcentajeTipoB float64 = 0.2
	PorcentajeTipoA float64 = 0.5
)

func calcularSalarioPorHoras(salario, horas float64) float64 {
	return salario * horas
}

func incrementoPorcentual(salario, incremento float64) float64 {
	return salario * incremento
}

func minutosAHoras(minutos int) float64 {
	return float64(minutos) / 60.0
}

func calcularSalario(minutos int, categoria string) float64 {
	salario := 0.0
	horas := minutosAHoras(minutos)
	switch categoria {
	case "C":
		salario = calcularSalarioPorHoras(C, horas)
		return salario
	case "B":
		salario = calcularSalarioPorHoras(B, horas)
		return salario + incrementoPorcentual(salario, PorcentajeTipoB)
	case "A":
		salario = calcularSalarioPorHoras(A, horas)
		return salario + incrementoPorcentual(salario, PorcentajeTipoA)
	}
	return salario
}

func main() {

	var minutos int
	var categoria string
	fmt.Println("Ingresa la cantidad de minutos")
	fmt.Scanf("%d", &minutos)

	fmt.Println("Ingresa la categoria")
	fmt.Scanf("%s", &categoria)

	var respuesta = calcularSalario(minutos, categoria)

	fmt.Printf("El salario es: %f\n", respuesta)

}
