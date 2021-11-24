package main

import "fmt"

func main()  {
	
	fmt.Println("Hola")
	var edad uint
	var trabaja bool
	var sueldo uint
	var tiempoTrabajando uint

	edad = 22
	sueldo = 99000
	trabaja = true
	tiempoTrabajando = 2

	if edad >= 22 && trabaja && tiempoTrabajando >= 1{
		if sueldo > 100000 {
			fmt.Println("Clasificas para el prestamo sin intereses!!!")
		} else {
			fmt.Println("Clasificas para el prestamo pero pagando intereses :(")
		}
	} else {
		fmt.Println("No clasificas por estos motivos")
		if edad < 22 {
			fmt.Println("No cumplis con la edad")
		}
		if !trabaja {
			fmt.Println("No tenes trabajo")
		}
		if tiempoTrabajando < 1 {
			fmt.Println("Sos muy reciente en tu trabajo")
		}
	}

}