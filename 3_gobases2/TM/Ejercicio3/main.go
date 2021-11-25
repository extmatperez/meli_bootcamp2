package main

import "fmt"

func datos() int {
	var minutos int
	var categoria string
	var horas int
	var sueldo int
	fmt.Scanln("ingrese la cantidad de minutos", &minutos)
	fmt.Scanln("ingrese la categoria", &categoria)

	horas = minutos / 60
	switch {
	case categoria == "A":
		sueldo = horas * 3000
	case categoria == "B":
		sueldo = horas * 1500
	case categoria == "C":
		sueldo = horas * 1000
	}

	return sueldo
}

func main() {

}
