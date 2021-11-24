package main

import "fmt"

func main() {
	var personas []persona = crearPersonas()
	validarAccesibilidadPrestamo(personas)
}

/*
	Una solucion sin utilizar estructura, puede ser crear un mapa string slice (string uint32) (si el sueldo es 0, no tiene trabajo)
*/

type persona struct {
	Nombre            string
	Edad              uint8
	EstaEmpleado      bool
	Sueldo            float64
	AntiguedadLaboral uint8
}

func crearPersonas() []persona {
	persona1 := persona{"Ramiro Gonzalez", 14, false, 0, 0}
	persona2 := persona{"Juan Diaz", 56, true, 54787.00, 2}
	persona3 := persona{"Maria Marquez", 38, true, 59353.00, 1}
	persona4 := persona{"Josefina Antunez", 28, true, 19002.00, 5}
	persona5 := persona{"Alberto Dominguez", 64, true, 205872.00, 15}
	persona6 := persona{"Miguel Bustamante", 56, false, 0, 0}

	var personas = []persona{persona1, persona2, persona3, persona4, persona5, persona6}
	return personas
}

func validarAccesibilidadPrestamo(personas []persona) {
	for _, persona := range personas {
		switch {
		case persona.Edad > 22 && persona.AntiguedadLaboral > 1:
			if persona.Sueldo > 100000 {
				fmt.Printf("A %v no se le cobraran intereses, debido al monto de su sueldo \n", persona.Nombre)
			} else {
				fmt.Printf("A %v se le cobraran intereses \n", persona.Nombre)
			}
		default:
			fmt.Printf("A %v no se le otorga prestamo \n", persona.Nombre)
		}
	}
}
