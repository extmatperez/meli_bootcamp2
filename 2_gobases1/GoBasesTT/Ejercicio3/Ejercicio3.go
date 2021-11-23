package main

import "fmt"

func main() {
	var edad int = 24
	var empleado bool = true
	var antiguedad int = 2
	var sueldo float32 = 105000.00
	prestamoBancario(edad, empleado, antiguedad, sueldo)

	edad = 23
	empleado = true
	antiguedad = 2
	sueldo = 99000
	prestamoBancario(edad, empleado, antiguedad, sueldo)

	edad = 22
	empleado = true
	antiguedad = 2
	sueldo = 99000
	prestamoBancario(edad, empleado, antiguedad, sueldo)
}

func prestamoBancario(edad int, empleado bool, antiguedad int, sueldo float32) {

	if edad > 22 && empleado && antiguedad > 1 {
		if sueldo < 100000 {
			fmt.Println("Puede recibir el prestamo, pero se cobraran intereses")
		} else {
			fmt.Println("Puede recibir el prestamo, y no se le cobraran intereses")
		}
	} else {
		fmt.Println("No puede acceder a un prestamo por no cumplir con los requisitos basicos")
	}
}
