package ordenamiento

import (
	"sort"
)

func OrderingSlice(x []int) []int {
	sort.Ints(x)
	return x
}
