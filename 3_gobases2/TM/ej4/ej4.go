package ej4

import (
	"errors"
	"fmt"
)

const (
	min  = "Min"
	prom = "Prom"
	max  = "Max"
)

func Ej4(operation string, values ...float64) float64 {
	fmt.Println("")
	op, err := Op(operation)

	if err != nil {
		panic("Wrong operation!")
	}

	return op(values...)
}

func Op(opType string) (func(...float64) float64, error) {
	switch opType {
	case min:
		return Min, nil
	case prom:
		return Prom, nil
	case max:
		return Max, nil
	default:
		return nil, errors.New("operation not found")
	}
}

func Prom(values ...float64) float64 {
	result := 0.0

	for _, val := range values {
		result += val
	}
	fmt.Println(result / float64(len(values)))
	return result / float64(len(values))
}

func Min(values ...float64) float64 {
	min := values[0]

	for _, el := range values {
		if el < min {
			min = el
		}
	}
	fmt.Println(min)
	return min
}

func Max(values ...float64) float64 {
	max := values[0]

	for _, el := range values {
		if el > max {
			max = el
		}
	}
	fmt.Println(max)
	return max
}
