package main

import (
	"errors"
	"fmt"
)

func main() {
	avg, err := average(6,2,3,4)

	if err != nil {
		fmt.Println("No se puede enviar un numero negativo!")
	} else {
		fmt.Println("El promedio es: ", avg)
	}

}

func average(grades ...float64) (float64, error) {
	fmt.Println("los numeros son: ", grades)
	var (
		l float64 = 0
		total float64 = 0
		avg float64 = 0
	)

	for _, i := range grades {
		if i < 0 {
			return 0, errors.New("Hay un negativo")
		}
		l = l + 1
		total = total + i
	}
	avg = total / l
	return avg, nil
}