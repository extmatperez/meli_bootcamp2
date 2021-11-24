package main

import "fmt"

func calcSalario(horas int, categoria string) float64 {
	switch categoria {
	case "C":
		return categoriaC(horas)
	case "B":
		return categoriaB(horas)
	case "A":
		return categoriaA(horas)
	default:
		return 0.0
	}
}

func categoriaC(horas int) float64 {
	salario := 1000 * horas
	return float64(salario)
}

func categoriaB(horas int) float64 {
	var salario float64 = 1500 * float64(horas)
	porcentaje := salario * 0.2
	salario = salario + porcentaje
	return salario
}

func categoriaA(horas int) float64 {
	var salario float64 = 3000 * float64(horas)
	porcentaje := salario * 0.5
	salario = salario + porcentaje
	return salario
}

func main() {
	fmt.Printf("El salario es %v \n", calcSalario(200, "A"))
}
