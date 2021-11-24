package main

import "fmt"

func main() {

	var nombre string = "Damian"
	var edad int = 25
	var empleado bool = true
	antiguedad := 2
	sueldo := 99.000
	if edad < 22 {
		fmt.Printf("El empleado %v no cumple requisito de edad\n", nombre)
	} else if empleado == true && antiguedad >= 1 && sueldo >= 100.000 {
		fmt.Printf("El empleado %v podrá acceder al prèstamo y no se le cobrara interes\n", nombre)
	} else if empleado == true && antiguedad >= 1 && sueldo <= 100.000 {
		fmt.Printf("El empleado %v podrá acceder al prèstamo pero se le cobrar un interes\n", nombre)
	}

}
