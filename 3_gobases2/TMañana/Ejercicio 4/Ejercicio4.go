package main

import "fmt"
import "errors"

const (
	minimo = "minimo"
	promedio = "promedio"
	maximo = "maximo"
 )

 
func minima(calif ...int) int {
	min := 10
	for _, ca := range calif {
		if ca < min {
			min = ca
		}
	}
	return min
}


func maxima(calif ...int) int {
	max := 0
	for _, ca := range calif {
		if ca > max {
			max = ca
		}
	}
	return max
}


func prom(calif ...int) int {
	sum := 0
	for _, ca := range calif {
		sum += ca
	}
	return sum / len(calif)
}


func estadisticas(operacion string) (func(calif ...int) int, error) {
	switch operacion {
	case minimo:
		return minima, nil
	case maximo:
		return maxima, nil
	case promedio:
		return prom, nil
	}
	return nil, errors.New("Operacion invalida")
}


func main() {

	calc, _ := estadisticas(minimo)
	calc1, _ := estadisticas(maximo)
	calc2, _ := estadisticas(promedio)

	res := calc(5,9,7,10,7)
	res1 := calc1(5,9,7,10,7)
	res2 := calc2(5,9,7,10,7)

	fmt.Println("Min:", res, "Max:", res1,"Prom:", res2)

	_, err := estadisticas("media")

	if err != nil {
		fmt.Println("Ocurrió un error", err)
	}

}