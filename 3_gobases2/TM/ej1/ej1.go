package ej1

import "fmt"

func Ej1(salary float64) float64 {
	if salary > 150000 {
		fmt.Println(0.1)
		return 0.27
	} else {
		fmt.Println(0.17)
		return 0.17
	}
}
