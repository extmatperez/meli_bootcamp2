package main

import "fmt"

func main() {

	var edad int
	var trabajo string
	var antiguedad int
	var sueldo float64

	fmt.Println("¿Que edad tienes?")
	fmt.Scanln(&edad)

	fmt.Println("¿Esta trabajando actualmente responda si/no?")
	fmt.Scanln(&trabajo)

	fmt.Println("¿Que antiguedad tiene en la empresa, en años?")
	fmt.Scanln(&antiguedad)

	fmt.Println("¿Cuanto es su sueldo?")
	fmt.Scanln(&sueldo)

	if edad > 22 && trabajo == "si" && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Puedes acceder a un prestamo y no se te cobrara interes")
		} else {
			fmt.Println("Puedes acceder a un prestamo pero se te cobrara interes")
		}
	} else {
		fmt.Println("En este momento no puedes acceder a un prestamo")
	}
}
