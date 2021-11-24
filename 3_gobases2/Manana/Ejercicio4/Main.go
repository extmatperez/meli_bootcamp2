package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	maximo   = "maximo"
	promedio = "promedio"
)

func findMinimo(numeros ...int) float64 {
	var menor int
	menor = numeros[0]

	for _, numero := range numeros {
		if numero < menor {
			menor = numero
		}
	}

	return float64(menor)
}

func findMaximo(numeros ...int) float64 {
	var mayor int
	mayor = numeros[0]

	for _, numero := range numeros {
		if numero > mayor {
			mayor = numero
		}
	}

	return float64(mayor)
}

func findPromedio(numeros ...int) float64 {
	var total int
	total = 0
	for _, numero := range numeros {
		total += numero
	}

	return float64(total) / float64(len(numeros))
}

func proveedora(procedimiento string) (func(...int) float64, error) {
	switch procedimiento {
	case minimo:
		return findMinimo, nil
	case maximo:
		return findMaximo, nil
	case promedio:
		return findPromedio, nil
	default:
		return nil, errors.New("parametro no valido")
	}
}

func main() {

	fmt.Println("Bienvenidos al ejercicio 4", minimo, maximo, promedio)

	fmt.Println("CASO 1")
	funcResultado1, err := proveedora("minimo")
	if err == nil {
		resulFinal1 := funcResultado1(6, 7, 8, 9, 10, 1, 3, 0, 3)
		fmt.Println("El resultado final minimo es", resulFinal1)
	} else {
		fmt.Println(err)
	}

	fmt.Println("CASO 2")
	funcResultado2, err := proveedora("maximo")
	if err == nil {
		resulFinal2 := funcResultado2(6, 7, 8, 9, 10, 1, 3, 0, 3)
		fmt.Println("El resultado final maximo es", resulFinal2)
	} else {
		fmt.Println(err)
	}

	fmt.Println("CASO 3")
	funcResultado3, err := proveedora("promedio")
	if err == nil {
		resulFinal3 := funcResultado3(6, 7, 8, 9, 10, 1, 3, 0, 3)
		fmt.Println("El resultado final del promedio es", resulFinal3)
	} else {
		fmt.Println(err)
	}
}
