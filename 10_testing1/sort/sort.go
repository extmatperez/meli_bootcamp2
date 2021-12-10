package sort

import (
	"sort"
)

func MethodSort(num []int) []int {
	var sortedNum []int
	sort.Slice(num, func(i, j int) bool {
		return num[i] < num[j]
	})
	for _, v := range num {
		sortedNum = append(sortedNum, v)
	}
	return sortedNum
}
