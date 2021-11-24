package main

import (
	"fmt"
)

func salario(minutos int, categoria string) (total float64) {
	switch categoria {
	case "C":
		total = (float64(minutos) / 60) * 1000
	case "B":
		total = (float64(minutos) / 60) * 1500
		total += (0.2 * total)
	case "A":
		total = (float64(minutos) / 60) * 3000
		total += (0.5 * total)
	default:
		total = 0

	}
	return total
}

func main() {
	base := 600
	categoria := "A"
	fmt.Printf("Minutos trabajados :%v \nCategoria:%v \nEl total es de :%v\n", base, categoria, salario(base, categoria))

}
