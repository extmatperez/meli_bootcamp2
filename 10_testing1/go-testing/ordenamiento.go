package calculadora

import "sort"

func Ordenar(lista []int) []int {
	for i, value := range lista {
		for j := i + 1; j < len(lista); j++ {
			if value > lista[j] {
				lista[i], lista[j] = lista[j], value
			}
		}
	}
	return lista
}

func OrdenarSort(lista []int) {
	sort.Ints(lista)
}
