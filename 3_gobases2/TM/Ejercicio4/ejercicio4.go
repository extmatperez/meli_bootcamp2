package main

import (
	"errors"
	"fmt"
)

const (
	MIN = "minimo"
	AVG = "promedio"
	MAX = "maximo"
)

func main() {
	minFunc, errMin := operation(MIN)
	avgFunc, errAvg := operation(AVG)
	maxFunc, errMax := operation(MAX)

	if errMin != nil {
		fmt.Println(errMin)
	} else {
		minValue := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("La calificación mínima es:", minValue)
	}

	if errAvg != nil {
		fmt.Println(errAvg)
	} else {
		avgValue := avgFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("El promedio de calificaciones es:", avgValue)
	}

	if errMax != nil {
		fmt.Println(errMax)
	} else {
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("La calificación más alta es:", maxValue)
	}
}

func operation(function string) (func(scores ...int) float64, error) {
	switch function {
	case MIN:
		return minimum, nil
	case AVG:
		return average, nil
	case MAX:
		return maximum, nil
	default:
		errorMsg := "la operacion de '" + function + "' no esta definida"
		return func(scores ...int) float64 { return 0.00 }, errors.New(errorMsg)
	}
}

func minimum(scores ...int) float64 {
	var min int = 0

	for i := 0; i < len(scores); i++ {
		if i == 0 || scores[i] < min {
			min = scores[i]
		}
	}

	return float64(min)
}

func average(scores ...int) float64 {
	var sum int = 0

	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	return float64(sum) / float64(len(scores))
}

func maximum(scores ...int) float64 {
	var max int = 0

	for i := 0; i < len(scores); i++ {
		if i == 0 || scores[i] > max {
			max = scores[i]
		}
	}

	return float64(max)
}
