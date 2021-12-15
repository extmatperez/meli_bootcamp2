package main

import "fmt"

func main() {
	fmt.Println(calcularSalario(1000, "A"))
	fmt.Println(calcularSalario(1000, "B"))
	fmt.Println(calcularSalario(1000, "C"))
}

func calcularSalario(minutos float64, category string) float64 {
	horas := minutos / 60
	switch category {
	case "A":
		base := 3000 * horas
		return base + base*0.5
	case "B":
		base := 1500 * horas
		return base + base*0.2
	case "C":
		base := 1000 * horas
		return base
	default:
		return 0
	}
}
