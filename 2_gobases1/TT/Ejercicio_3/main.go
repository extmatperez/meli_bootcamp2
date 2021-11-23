package main

import "fmt"

func main() {
	edad := 25
	empleado := true
	mesesAntiguedad := 12
	sueldo := 800000

	if edad >= 22 && empleado && mesesAntiguedad >= 12 {
		if sueldo > 100000 {
			fmt.Println("Felicidades! Eres apto para recibir el préstamo pero debes pagar intereses")
		} else {
			fmt.Println("Felicidades! Eres apto para recibir el préstamo SIN pagar intereses")
		}
	} else {
		fmt.Println("Lo sentimos! No eres apto para recibir el préstamo")
	}
}
