package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {

	minFunc, err := operacion(minimo)
	promFunc, errorProm := operacion(promedio)
	maxFunc, errorMax := operacion(maximo)

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Valor Minimo: %d", valorMinimo)
	}
	if errorProm != nil {
		fmt.Println(errorProm)
	} else {
		fmt.Printf("\n valor Promedio: %d", valorPromedio)
	}
	if errorMax != nil {
		fmt.Println(errorMax)
	} else {
		fmt.Printf("\nValor Maximo: %d", valorMaximo)
	}

}

func calcularMin(valores ...int) (min int) {
	if len(valores) > 0 {
		min = valores[0]
	}
	for i := 1; i < len(valores); i++ {
		if valores[i] < min {
			min = valores[i]
		}
	}
	return
}

func calcularPromedio(valores ...int) int {
	valoresSum := 0
	for _, value := range valores {
		valoresSum = valoresSum + value
	}
	return valoresSum / len(valores)
}

func calcularMax(valores ...int) (max int) {
	if len(valores) > 0 {
		max = valores[0]
	}
	for i := 1; i < len(valores); i++ {
		if valores[i] > max {
			max = valores[i]
		}
	}
	return
}

func operacion(tipoDeCalculo string) (func(valores ...int) int, error) {

	switch tipoDeCalculo {
	case "minimo":
		return calcularMin, nil
	case "promedio":
		return calcularPromedio, nil
	case "maximo":
		return calcularMax, nil
	}

	return nil, errors.New("Error")
}
