package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {

	oper, err := operacion(minimo)
	if err == nil {
		fmt.Printf("Minimo: %.0f\n", oper(4, 7, 2, 9, 3, 3, 5))
	} else {
		fmt.Printf("Error: %s.\n", err)
	}
	oper, err = operacion(promedio)
	if err == nil {
		fmt.Printf("Promedio: %.2f\n", oper(4, 7, 2, 9, 3, 3, 5))
	} else {
		fmt.Printf("Error: %s.\n", err)
	}
	oper, err = operacion(maximo)
	if err == nil {
		fmt.Printf("Maximo: %.0f\n", oper(4, 7, 2, 9, 3, 3, 5))
	} else {
		fmt.Printf("Error: %s.\n", err)
	}
	oper, err = operacion("otro")
	if err == nil {
		fmt.Printf("Otro: %.2f.\n", oper(4, 7, 2, 9, 3, 3, 5))
	} else {
		fmt.Printf("Error: %s.\n", err)
	}
}

func CalcularMinimo(notas ...int) float64 {
	minimo := math.MaxInt
	for _, n := range notas {
		if n < minimo {
			minimo = n
		}
	}
	return float64(minimo)
}

func CalcularMaximo(notas ...int) float64 {
	maximo := -1
	for _, n := range notas {
		if n > maximo {
			maximo = n
		}
	}
	return float64(maximo)
}

func CalcularPromedio(notas ...int) (promedio float64) {
	promedio = 0
	for _, n := range notas {
		promedio += float64(n)
	}
	promedio /= float64(len(notas))
	return
}

func operacion(tipo string) (func(...int) float64, error) {
	switch tipo {
	case "minimo":
		return CalcularMinimo, nil
	case "promedio":
		return CalcularPromedio, nil
	case "maximo":
		return CalcularMaximo, nil
	default:
		return nil, errors.New("Operacion inválida")
	}
}
