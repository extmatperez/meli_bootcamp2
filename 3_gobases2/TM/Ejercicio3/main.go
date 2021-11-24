package main

import "fmt"

func main() {
	fmt.Printf("El salario es: %.3f", salario(10000, "B"))

}

func salario(minutos int, categoria string) float64 {
	resultado := 0.0
	horas := float64(minutos) / 60.0

	if categoria == "C" {
		resultado = horas * 1000
	} else if categoria == "B" {
		var mensual = float64(horas * 1500.0)
		resultado = float64(mensual) + float64(mensual*(0.2))
	} else {
		var mensual = (horas * 3000)
		resultado = mensual + (mensual * (0.5))
	}

	return resultado
}
