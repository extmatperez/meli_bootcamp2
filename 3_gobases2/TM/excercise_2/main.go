package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(CalcPromedio(5, 5, 3))
}

func CalcPromedio(notas ...float64) (float64, error) {
	var promedio float64
	var cantNotas int
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("hay un numero negativo")

		}
		promedio += nota
		cantNotas++
	}
	promedio = promedio / float64(cantNotas)
	return promedio, nil
}
