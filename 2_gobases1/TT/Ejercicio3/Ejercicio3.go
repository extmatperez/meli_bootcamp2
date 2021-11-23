package main

import (
	"fmt"
)

func main() {
	var years, work_exp int
	var salary float64
	fmt.Println("¿Cuantos años tiene?")
	fmt.Scanf("%d", &years)
	fmt.Println("¿Hace cuantos años trabaja?")
	fmt.Scanf("%d", &work_exp)
	fmt.Println("¿Cual es su salario?")
	fmt.Scanf("%f", &salary)

	if years > 22 {
		if work_exp > 1 {
			if salary > 100000 {
				fmt.Println("Puede recibir el prestamo y sin intereses por su alto salario.")
			} else {
				fmt.Println("Puede recibir el prestamo pero con intereses porque su sueldo no supera los $100k")
			}
		} else {
			fmt.Println("No puede recibir el prestamo dado que no cumple con los requisitos.")
		}
	} else {
		fmt.Println("No puede recibir el prestamo dado que no cumple con los requisitos.")
	}
}
