package go_testing

import "sort"

func IntArrayOrder(ints []int) []int {
	var auxiliar int
	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints); j++ {
			if ints[i] < ints[j] {
				auxiliar = ints[i]
				ints[i] = ints[j]
				ints[j] = auxiliar
			}
		}
	}
	return ints
}

func orderSlices (ints []int) []int {
	sort.Ints(ints)
	return ints
}
