package ordenar

import "sort"

func Ordenar(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a
}
