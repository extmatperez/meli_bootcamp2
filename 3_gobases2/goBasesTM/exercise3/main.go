package main

import "fmt"

func main() {
	var categoria string
	var minutos int
	fmt.Println("Ejercicio 3")
	fmt.Println("Por favor ingresa la categoria del empleado")
	fmt.Scanf("%s", &categoria)
	fmt.Println("Por favor ingresa los minutos trabajados del empleado")
	fmt.Scanf("%d", &minutos)
	fmt.Println(salario(minutos, categoria))

}

func salario(minutos int, categoria string) int {

	var salario int
	var horas int = minutos / 60

	if categoria == "a" || categoria == "A" {
		salario = ((horas * 3000) * 15) / 10
	} else if categoria == "b" || categoria == "B" {
		salario = ((horas * 1500) * 12) / 10
	} else if categoria == "c" || categoria == "C" {
		salario = horas * 1000
	} else {
		fmt.Println("La categoria que usaste no existe volve a ejecutar el programa por favor")
	}

	return salario
}
