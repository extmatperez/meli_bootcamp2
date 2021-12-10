package main

import "fmt"

func main() {

	estudiante := []string{"Benjamin", "Nahuel", "Brenda", "Marcos",
		"Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán",
		"Leandro", "Eduardo", "Duvraschka"}

	fmt.Println("Lista inicial", estudiante)

	estudiante = append(estudiante, "Gabriela")

	fmt.Println("Lista final", estudiante)

}
