package main

import "fmt"

func main() {

	var edad int = 22
	var antiguedad float64 = 1
	var sueldo int = 90000

	if edad < 22 {
		fmt.Println("No cumple con la edad minima para acceder al credito")
	} else if antiguedad <= 1 {
		fmt.Println("No cumple con la antiguedad minima requerida")
	} else {
		if sueldo > 100000 {
			fmt.Println("Su sueldo es superior a $100.000, por lo que no se le cobrara interes")
		} else {
			fmt.Println("Su sueldo es menor a $100.000 por lo que se le cobrara interes")
		}
	}
}
