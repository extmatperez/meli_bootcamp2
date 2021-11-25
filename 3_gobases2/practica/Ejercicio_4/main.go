package main

import "fmt"

func main() {
	fmt.Println(operator("maximo", 2, 3, 2, 1, 2, 3, 4, 5, 6))
}

func operator(key string, values ...int) int {
	switch key {
	case "minimo":
		return caller(values, minimum)
	case "maximo":
		return caller(values, maximum)
	case "promedio":
		return caller(values, average)
	}
	return 0
}
func caller(values []int, operation func(numbers ...int) int)  int {
	var result int
	for i, valor := range values {
		if i == 0 {
			result = valor
		} else {
			result = operation(result, valor)
		} }
	return result
}
func minimum(numbers ...int) int {
	var (
		mini int
	)
	for i, checker := range numbers{
		if i==0 || checker < mini {
			mini = checker
		}
	}
	return mini
}
func maximum(numbers ...int) int {
	var (
		maxi int
	)
	for i, checker := range numbers{
		if i==0 || checker > maxi {
			maxi = checker
		}
	}
	return maxi
}
func average(numbers ...int) int {
	var (
		avg int
		count int
	)
	for _, checker := range numbers{
		count = count + checker
	}
	avg = count / len(numbers)
	return avg
}