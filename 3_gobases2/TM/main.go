package main

import (
	"fmt"

	"github.com/lean1097/meli_bootcamp2/3_gobases2/TM/ej1"
	"github.com/lean1097/meli_bootcamp2/3_gobases2/TM/ej2"
	"github.com/lean1097/meli_bootcamp2/3_gobases2/TM/ej3"
	"github.com/lean1097/meli_bootcamp2/3_gobases2/TM/ej4"
	"github.com/lean1097/meli_bootcamp2/3_gobases2/TM/ej5"
)

func main() {
	fmt.Println("Go Bases TM")

	fmt.Println("Ejercicio 1")
	ej1.Ej1(555)
	ej1.Ej1(155500)

	fmt.Println("Ejercicio 2")
	ej2.Ej2(1, 2, 3, 4, 5, 6, 7)
	ej2.Ej2(1, 2, 3, 4, -5, 6, 7)

	fmt.Println("Ejercicio 3")
	ej3.Ej3("C", 100)
	ej3.Ej3("B", 300)
	ej3.Ej3("A", 300)

	fmt.Println("Ejercicio 4")
	ej4.Ej4("Min", 1, 5, 10)
	ej4.Ej4("Prom", 8, 10, 9)
	ej4.Ej4("Max", 1, 5, 10)

	fmt.Println("Ejercicio 5")
	ej5.Ej5("Tar", 5)
	ej5.Ej5("Ham", 5)
	ej5.Ej5("Perro", 2)
	ej5.Ej5("Gato", 4)
}
