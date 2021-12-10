package gotesting

import (
	"fmt"
	"sort"
)

func Restar(a, b int) int {
	return a - b
}

func Ordenar(a []int) []int {

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	return a
}

func Dividir(numerador, denominador int) (int, error) {

	if denominador == 0 {
		return 0, fmt.Errorf("El denominador no puede ser 0 %v", denominador)
	} else {
		return numerador / denominador, nil
	}

}
