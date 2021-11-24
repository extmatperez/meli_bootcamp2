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

	minFunc, err := op(minimo)
	promFunc, err := op(promedio)
	maxFunc, err := op(maximo)
	undefinedFunc, err := op("sarasa")

	valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorPromedio := promFunc(2, 3, 3, 4, 1, 2, 4, 5)
	valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(valorMinimo)
	fmt.Println(valorPromedio)
	fmt.Println(valorMaximo)
	if err != nil {
		fmt.Println("También hubo un error")
	} else {
		fmt.Printf("en lugar de error, se retornó %v", undefinedFunc)
	}

}

func op(operation string) (func(m ...int) int, error) {

	switch operation {
	case minimo:
		return min, nil
	case maximo:
		return max, nil
	case promedio:
		return avg, nil
	default:
		return nil, errors.New("undefined operation")
	}

}

func min(marks ...int) int {

	currentMin := marks[0]

	for i := 1; i < len(marks); i++ {
		if marks[i] < currentMin {
			currentMin = marks[i]
		}
	}
	return currentMin
}

func max(marks ...int) int {

	currentMax := marks[0]

	for i := 1; i < len(marks); i++ {
		if marks[i] > currentMax {
			currentMax = marks[i]
		}
	}
	return currentMax
}

func avg(marks ...int) int {

	currentSum := 0

	for _, num := range marks {
		currentSum += num
	}

	return currentSum / len(marks)
}
