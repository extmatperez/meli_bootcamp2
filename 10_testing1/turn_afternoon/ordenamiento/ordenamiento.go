package ordenamiento

import "sort"

func OrdenarNumeros(numeros ...int) []int {

	var newNumeros []int
	for _, n := range numeros {
		newNumeros = append(newNumeros, n)
	}

	sort.Ints(newNumeros)

	return newNumeros
}
