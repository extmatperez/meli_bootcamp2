package calculadora

import "sort"

func Ordenar(slice []int) []int {
	sort.Ints(slice)
	return slice
}
