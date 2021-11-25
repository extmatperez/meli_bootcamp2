package main

import (
	"fmt"
)

type fecha struct {
	dia int
	mes int
	año int
}

type Alumno struct {
	nombre   string
	apellido string
	dni      int
	ingreso  fecha
}

func (fec fecha) detalleFecha() {
	fmt.Printf("%v/%v/%v\n", fec.dia, fec.mes, fec.año)
}

func (alu Alumno) detalle() {
	var fec fecha = alu.ingreso
	fmt.Printf("Nombre: %v\n", alu.nombre)
	fmt.Printf("Apellido: %v\n", alu.apellido)
	fmt.Printf("DNI: %v\n", alu.dni)
	fmt.Printf("Fecha: ")
	fec.detalleFecha()
}

func main() {
	alu := Alumno{
		nombre:   "Pablo",
		apellido: "Perez",
		dni:      315485478,
		ingreso: fecha{
			dia: 19,
			mes: 8,
			año: 1975,
		},
	}

	alu.detalle()
}
