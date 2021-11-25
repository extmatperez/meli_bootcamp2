package main

import "fmt"

func main() {
	var categoria string
	var minutos float64
	fmt.Println("ingrese categoria")
	fmt.Scanln(&categoria)
	fmt.Println("ingrese los minutos laboradas")
	fmt.Scanln(&minutos)
	horas := minutos / 60
	fmt.Println("El salario correspondiente al empleado es:", CalSalario(categoria, horas))
}

func CalSalario(categoria string, horas float64) float64 {
	switch categoria {
	case "A":
		return (3000 * horas) + (3000 * horas * 0.5)
	case "B":
		return (1500 * horas) + (1500 * horas * 0.2)
	case "C":
		return 1000 * horas
	default:
		return 0
	}

}
