package main

import "fmt"

func main() {
	fmt.Println("El impuesto de 75000 es: ", impuestoDeSalario(75000))

	fmt.Println("El impuesto de 175000 es: ", impuestoDeSalario(175000))

	fmt.Println("El impuesto de 10000 es: ", impuestoDeSalario(10000))

}

func impuestoDeSalario(salario float64) float64 {
	descuento := 0

	if salario > 50000 {
		descuento = 17
	} else if salario > 150000 {
		descuento = 10
	}

	return salario * float64(descuento) / 100
}
