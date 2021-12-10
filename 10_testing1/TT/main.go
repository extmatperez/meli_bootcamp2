package main

import (
	"errors"
	"sort"
)

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Ordenar(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}

func Dividir(num, dem float64) (float64, error) {
	if dem == 0 {
		return 0, errors.New("no se puede dividir por 0")
	}
	return num / dem, nil
}
