package main

import "fmt"

func calcularSalario(categoria string, minutosTrabajados int) float64 {
	horasTrabajadas := float64(minutosTrabajados) / 60
	switch categoria {
	case "C":
		return horasTrabajadas * 1000
	case "B":
		sueldoMes := horasTrabajadas * 1500
		return sueldoMes + ((sueldoMes * 20) / 100)
	case "A":
		sueldoMes := horasTrabajadas * 3000
		return sueldoMes + ((sueldoMes * 50) / 100)

	}
	return 0
}

func main() {

	categoria := "A"
	minutosTrabajados := 60

	fmt.Println("El salario mensual es de : ", calcularSalario(categoria, minutosTrabajados))
}
