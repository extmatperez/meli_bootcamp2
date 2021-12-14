package ordenamiento

import "sort"

func AscendingOrder(ns []int) []int {
	sort.Ints(ns)
	return ns
}
