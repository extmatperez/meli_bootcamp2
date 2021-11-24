package main

import "fmt"

func calculaImpuestos(salario float32) int {

	var porcentaje int

	if salario > 150000 {
		porcentaje = 27
	} else if salario > 50000 {
		porcentaje = 17
	} else {
		porcentaje = 0
	}

	return porcentaje

}

func sueldoRestante(sueldo int, impuestos int) int {
	var restante = sueldo - sueldo*impuestos/100
	return restante
}

func main() {

	fmt.Println("Bienvenidos al ejercicio 1")

	fmt.Println("Caso 1")
	salarioEmpleado1 := 135000
	respuesat1 := calculaImpuestos(float32(salarioEmpleado1))
	fmt.Println("El empleado 1 tiene en impuestos", respuesat1, "esto seria un total de", sueldoRestante(salarioEmpleado1, respuesat1))

	fmt.Println("Caso 2")
	salarioEmpleado2 := 199000
	respuesat2 := calculaImpuestos(float32(salarioEmpleado2))
	fmt.Println("El empleado 2 tiene en impuestos", respuesat2, "esto seria un total de", sueldoRestante(salarioEmpleado2, respuesat2))

	fmt.Println("Caso 3")
	salarioEmpleado3 := 24000
	respuesat3 := calculaImpuestos(float32(salarioEmpleado3))
	fmt.Println("El empleado 3 tiene en impuestos", respuesat3, "esto seria un total de", sueldoRestante(salarioEmpleado3, respuesat3))

}
