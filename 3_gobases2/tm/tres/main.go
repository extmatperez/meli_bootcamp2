package main

import "fmt"

func calcularSalario(categoria string, minutos float64) float64 {
	salario := 0.0
	switch categoria {
	case "C":
		salario = 1000 * minutos / 60
	case "B":
		salario = 1500*minutos/60 + 1500*minutos/60*20/100
	case "A":
		salario = 3000*minutos/60 + 3000*minutos/60*50/100
	}
	return salario
}

func main() {
	fmt.Println(calcularSalario("B", 15600))
}
