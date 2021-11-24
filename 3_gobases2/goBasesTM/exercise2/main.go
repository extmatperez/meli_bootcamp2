package main

import (
	"errors"
	"fmt"
)

func main() {

	/* var values []int

	fmt.Println("Ejercicio 4")
	fmt.Println("Ingrese las notas que saco")

	fmt.Scanf("%s", &values)
	fmt.Println(promedio(values)) */
	fmt.Println(promedio(1, 2, 3, 4, 5, 6, 7))
}

func promedio(values ...int) (int, error) {

	var resultado int
	var promedio int

	for _, value := range values {
		fmt.Println(value)
		resultado += value
		if value < 0 {
			return 0, errors.New("no se pueden tener notas negativas")
		}
	}
	promedio = resultado / len(values)
	return promedio, nil
}
