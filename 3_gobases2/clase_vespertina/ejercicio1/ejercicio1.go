package main

import (
	"fmt"
)

func main() {
	empleados := crearEmpleados()
	calcularImpuestos(empleados)
}

type empleado struct {
	Nombre string
	Sueldo float64
}

func crearEmpleados() []empleado {
	empleado1 := empleado{"Ramiro Gonzalez", 56000}
	empleado2 := empleado{"Francisco Mitre", 86550}
	empleado3 := empleado{"Andres Otero", 16000}
	empleado4 := empleado{"Maria Gimenez", 103500}
	empleado5 := empleado{"Natalia Antonaz", 8000}
	empleado6 := empleado{"Virginia Tojo", 300500}
	empleado7 := empleado{"Tomas Montero", 50300}
	empleado8 := empleado{"Nahuel Bengochea", 49000}
	var empleados = []empleado{empleado1, empleado2, empleado3, empleado4, empleado5, empleado6, empleado7, empleado8}
	return empleados
}

func calcularImpuestos(empleados []empleado) {
	for _, emple := range empleados {
		porcentajeDescuento := obtenerDescuento(emple.Sueldo)
		totalDescuento := (porcentajeDescuento * emple.Sueldo) / 100
		fmt.Printf("%v tiene un sueldo de %v, el descuento correspondiente es de %v, llegando a un monto de descuento de %v \n", emple.Nombre, emple.Sueldo, porcentajeDescuento, totalDescuento)
	}
}

func obtenerDescuento(sueldo float64) float64 {
	descuentoBase := 17.0
	if sueldo > 50000 {
		if sueldo > 150000 {
			return descuentoBase + 10
		}
		return descuentoBase
	} else {
		return 0
	}
}
