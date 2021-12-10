package main

import (
	"fmt"
	"sort"
)

func ordering(x []int) []int {
	sort.Ints(x)
	return x
}

func main() {
	s := []int{1, 8, 3, 10, 58, 2}
	fmt.Println(ordering(s))
}