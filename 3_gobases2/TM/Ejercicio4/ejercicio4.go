package main

import "fmt"

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {
	values := []float64{10, 11, 12, 8, 9, 6, 10, 11}
	oper, err := operacion(maximo)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%.2f", oper(values...))
}

func operacion(operacion string) (func(...float64) float64, error) {
	switch operacion {
	case minimo:
		return minFunc, nil
	case promedio:
		return promFunc, nil
	case maximo:
		return maxFunc, nil
	default:
		return nil, fmt.Errorf("%v", "La operacion proporcionada no se encuentra contemplada")
	}
}

func minFunc(numeros ...float64) float64 {
	min := 15098579854.00
	for _, n := range numeros {
		if n < min {
			min = n
		}
	}
	return min
}

func promFunc(numeros ...float64) float64 {
	promedio := 0.00
	for _, n := range numeros {
		promedio += (float64(n) / float64((len(numeros))))
	}
	return promedio
}

func maxFunc(numeros ...float64) float64 {
	max := 0.00
	for _, n := range numeros {
		if n > max {
			max = n
		}
	}
	return max
}
