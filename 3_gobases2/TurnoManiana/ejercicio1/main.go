package main

import "fmt"

func impuestoSueldo(sueldo float64) float64 {
	if sueldo >= 150000 {
		return (sueldo * 27) / 100
	} else if sueldo >= 50000 && sueldo < 150000 {
		return (sueldo * 10) / 100
	} else {
		return 0.0
	}

}

func main() {

	sueldo := 49000.00
	impuestoSalario := impuestoSueldo(sueldo)
	fmt.Printf("El descuento sera de $%v \n", impuestoSalario)

}
