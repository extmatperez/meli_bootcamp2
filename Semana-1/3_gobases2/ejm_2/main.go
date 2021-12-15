package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio, err := calcularPromedioNotas(10, 20, -30, 40, 50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(promedio)
	}
}

func calcularPromedioNotas(notas ...float64) (float64, error) {
	if len(notas) == 0 {
		return 0, errors.New("no hay notas")
	}
	var suma float64 = 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("nota no válida")
		}
		suma += nota
	}
	return suma / float64(len(notas)), nil
}
