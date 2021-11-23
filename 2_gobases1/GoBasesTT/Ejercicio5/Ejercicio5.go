package main

import "fmt"

func main() {
	listadoDeNombres()
}

func listadoDeNombres() {
	var listadoNombres = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println(listadoNombres)

	listadoNombres = append(listadoNombres, "Gabriela")
	fmt.Println(listadoNombres)

}
