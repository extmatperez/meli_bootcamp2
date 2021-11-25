package main

import "fmt"

func datos() float64 {
	var minutos float64
	var categoria string
	var horas float64
	var sueldo float64
	fmt.Scanln(&minutos)
	fmt.Scanln(&categoria)

	horas = minutos / 60
	switch {
	case categoria == "A":
		sueldo = (horas * 3000) * 1.5
	case categoria == "B":
		sueldo = (horas * 1500) * 1.2
	case categoria == "C":
		sueldo = (horas * 1000)
	}

	return sueldo
}

func main() {
	fmt.Println(datos())
}
