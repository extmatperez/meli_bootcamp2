package main

import "fmt"

type Categorias int

const (
	A Categorias = iota
	B
	C
)

type Empleado struct {
	Categoria Categorias
	Minutos   int
}

func main() {
	e1 := Empleado{A, 130}
	e2 := Empleado{B, 130}
	e3 := Empleado{C, 130}
	fmt.Printf("El salario del empleado 1 es: %.2f\n", CalcularSalario(e1.Minutos, e1.Categoria))
	fmt.Printf("El salario del empleado 2 es: %.2f\n", CalcularSalario(e2.Minutos, e2.Categoria))
	fmt.Printf("El salario del empleado 3 es: %.2f\n", CalcularSalario(e3.Minutos, e3.Categoria))

}

func CalcularSalario(minutos int, cat Categorias) float64 {
	horas := minutos / 60
	switch cat {
	case A:
		return float64(3000*horas) * 1.5
	case B:
		return float64(1500*horas) * 1.2
	case C:
		return float64(1000 * horas)
	default:
		return 0.0
	}
}
