package main

import "fmt"

var listadoNombres = []string{
	"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro",
	"Axel", "Alez", "Dolores", "Federico", "Hernan",
	"Leandro", "Eduardo", "Duvraschka",
}

func agregarNombres(nombre string) {
	var listadoNombres = append(listadoNombres, nombre)
	fmt.Printf("Se sumo el estudiante %v, ahora el listado queda %v", nombre, listadoNombres)
}

func main() {
	agregarNombres("Gabriela")
}
