package main

import (
	"fmt"

	"github.com/lean1097/meli_bootcamp2/2_gobases1/TT/ej1"
	"github.com/lean1097/meli_bootcamp2/2_gobases1/TT/ej2"
	"github.com/lean1097/meli_bootcamp2/2_gobases1/TT/ej3"
	"github.com/lean1097/meli_bootcamp2/2_gobases1/TT/ej4"
	"github.com/lean1097/meli_bootcamp2/2_gobases1/TT/ej5"
	"github.com/lean1097/meli_bootcamp2/2_gobases1/TT/ej6"
)

func main() {
	fmt.Println("Go bases TT")

	fmt.Println("Ejercicio 1")
	ej1.Ej1("Leandro")

	fmt.Println("Ejercicio 2")
	ej2.Ej2(30.5, 0.10)

	fmt.Println("Ejercicio 3")
	ej3.Ej3(24, 1, 175000, true)
	ej3.Ej3(24, 1, 175000, false)
	ej3.Ej3(24, 1, 99000, true)

	fmt.Println("Ejercicio 4")
	ej4.Ej4(1)
	ej4.Ej4(5)
	ej4.Ej4(10)

	fmt.Println("Ejercicio 5")
	ej5.Ej5()
	ej5.AddStudent("Julian")

	fmt.Println("Ejercicio 6")
	ej6.Ej6()
	ej6.SearchEmployee("Benjamin")
	ej6.GreaterThan21()
	ej6.AddEmployee("Federico", 25)
	ej6.DeleteEmployee("Pedro")
}
