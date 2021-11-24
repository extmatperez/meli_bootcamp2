package main

import (
	pkg "github.com/extmatperez/meli_bootcamp2/blob/ghione_andres/go.mod"
)

func main() {
	alumnos := pkg.CargarAlumnos()
	for _, alumn := range alumnos {
		alumn.Detalle()
	}
}
