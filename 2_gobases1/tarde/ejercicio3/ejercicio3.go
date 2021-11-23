package main

import "fmt"

func prestamo(cliente map[string]int) {
	if cliente["edad"] > 22 && cliente["esEmpleado"] == 1 && cliente["antiguedad"] > 1 {
		if cliente["sueldo"] > 100000 {
			fmt.Println("Se le otorga un prestamo SIN intereses")
		} else {
			fmt.Println("Se le otorga un prestamo CON intereses")
		}
	} else {
		fmt.Println("NO se le puede otorgar el prestamo :C")
	}
}

var ivan = map[string]int{
	"edad":       23,
	"esEmpleado": 1,
	"sueldo":     50000,
	"antiguedad": 2,
}

func main() {
	prestamo(ivan)
}
