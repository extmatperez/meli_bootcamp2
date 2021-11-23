package main 

import "fmt"

func main() {
	edad := 25
	empleado := true
	antiguedad := 5
	sueldo := 100000

	if edad > 22 && empleado == true && antiguedad > 1 {
		if sueldo >= 100000 {
			fmt.Println("Usted aplica para un préstamo sin intereses!")
		} else {
			fmt.Println("Usted aplica para un préstamo!")
		}
	} else {
		fmt.Println("Usted no aplica para un prestamo")
	}
}