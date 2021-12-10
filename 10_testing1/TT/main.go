package main

import "sort"

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Ordenar(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}
