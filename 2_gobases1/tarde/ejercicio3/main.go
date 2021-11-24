package main

import "fmt"

func main() {

	var (
		edad       int  = 22
		empleado   bool = true
		antiguedad int  = 2
		salario    int  = 100001
	)

	if edad > 22 && empleado && antiguedad > 1 {
		fmt.Println("Apto para prestamo")
		if salario > 100000 {
			fmt.Println("No paga interes")
		}
	} else {
		fmt.Println("No apto para prestamo")
	}

}
