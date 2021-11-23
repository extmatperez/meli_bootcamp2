package main

import (
	"fmt"
)

func main() {
	edad := 50
	empleado := true
	tiempoTrabajo := 2
	sueldo := 100001
	fmt.Printf("edad: %v \nesta empleado:%v \ntiempo trabajando:%v \nsueldo: %v", edad, empleado, tiempoTrabajo, sueldo)

	if empleado && edad > 22 && tiempoTrabajo > 1 {
		fmt.Println("\n\nSI es posible otorgar prestamo")
		if sueldo > 100000 {
			fmt.Println("NO cobrara interes")
		} else {
			fmt.Println("SI cobrara interes")
		}
	} else {
		fmt.Println("No es posible otorgar prestamo")
	}

}
