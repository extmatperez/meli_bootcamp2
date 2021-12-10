package funciones

import (
	"errors"
	"sort"
)

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Ordenar(s []int) []int {
	sort.Ints(s)
	return s
}

func Dividir(num1, num2 int) (float64, error) {

	num1B := float64(num1)
	num2B := float64(num2)

	if num2 == 0 {
		return 0, errors.New("no se puede dividir en 0")
	}

	return num1B / num2B, nil
}
