package main

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
	dni      string
	fecha    string
}

func (p persona) detalle() {
	fmt.Printf("Nombre: %v\n", p.nombre)
	fmt.Printf("Apellido: %v\n", p.apellido)
	fmt.Printf("DNI: %v\n", p.dni)
	fmt.Printf("Fecha: %v\n", p.fecha)
}
func main() {
	p1 := persona{
		nombre:   "Diego",
		apellido: "Parra",
		dni:      "12345",
		fecha:    "16-09.1996",
	}
	fmt.Printf("Llamado función detalle:")
	p1.detalle()
}
