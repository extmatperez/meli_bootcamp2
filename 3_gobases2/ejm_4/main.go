package main

import (
	"errors"
	"fmt"
)

const (
	minimo = "minimo"
	promedio = "promedio"
	maximo = "maximo"
 )

func main() {
	minFunc, _ := operacion(minimo)
	promFunc, _ := operacion(promedio)
	maxFunc, _ := operacion(maximo)
	 
	valorMinimo := minFunc(2,3,3,4,1,2,4,5)
	valorPromedio := promFunc(2,3,3,4,1,2,4,5)
	valorMaximo := maxFunc(2,3,3,4,1,2,4,5)

	fmt.Println("Valor minimo:", valorMinimo)
	fmt.Println("Valor promedio:", valorPromedio)
	fmt.Println("Valor maximo:", valorMaximo)
}

func operacion(oper string) (func (nums ...float64) float64, error) {
	switch oper {
		case minimo:
			return min, nil
		case promedio:
			return prom, nil
		case maximo:
			return max, nil
		default:
			return nil, errors.New("operacion no valida")
	}
}

func min(nums ...float64) float64 {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func max(nums ...float64) float64 {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func prom(nums ...float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}
 