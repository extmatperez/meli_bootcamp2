package ej2

import "fmt"

func Ej2(price, discount float64) float64 {
	result := price - (price * discount)
	fmt.Println(result)
	return result
}
