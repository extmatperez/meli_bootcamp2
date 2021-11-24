package main

import "fmt"

func salario(minutos int, categoria string) {
	/*
		Una empresa marinera necesita calcular el salario de sus empleados
		basándose en la cantidad de horas trabajadas por mes y la categoría.
		Si es categoría C, su salario es de $1.000 por hora
		Si es categoría B, su salario es de $1.500 por hora más un %20 de
		su salario mensual
		Si es de categoría A, su salario es de $3.000 por hora más un %50
		de su salario mensual
		Se solicita generar una función que reciba por parámetro la
		cantidad de minutos trabajados por mes y la categoría, y que
		devuelva su salario.
	*/
	switch categoria {
	case "A":
		fmt.Printf("Salario = $%v\n", float64(minutos*4500/60))
	case "B":
		fmt.Printf("Salario = $%v\n", float64(minutos*1800/60))
	case "C":
		fmt.Printf("Salario = $%v\n", float64(minutos*1000/60))
	default:
		fmt.Printf("Categoria erronea\n")
	}
}

func main() {
	salario(10000, "A")
	salario(10000, "B")
	salario(10000, "C")
}
