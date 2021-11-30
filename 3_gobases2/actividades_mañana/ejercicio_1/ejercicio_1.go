package main

import "fmt"

func impuesto(sueldo float64) float64 {
	valor := 0.0
	if sueldo > 50000 && sueldo < 150000 {
		valor = sueldo * 0.17

	} else if sueldo > 150000 {
		valor = sueldo * 0.10
	}
	return valor

}

func main() {
	var salary float64
	fmt.Println("Salario: ")
	fmt.Scanf("%f", &salary)
	fmt.Println("Los impuestos correspondientes son de ", impuesto(salary))

}
