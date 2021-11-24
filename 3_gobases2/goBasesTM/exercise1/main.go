package main

import "fmt"

func main() {
	fmt.Println(impuesto())
}

func impuesto() int {
	var salario int
	var impuesto float64
	var retencion float64

	fmt.Println("Ejercicio 1")
	fmt.Println("Por favor ingrese su salario mensual")
	fmt.Scanf("%d", &salario)

	if salario >= 50000 && salario <= 150000 {
		impuesto = 0.17
		retencion = float64(salario) * impuesto
		return int(retencion)
	} else if salario >= 150000 {
		impuesto = 0.27
		retencion = float64(salario) * impuesto
		return int(retencion)
	} else {
		retencion = 0
		return int(retencion)
	}

}
