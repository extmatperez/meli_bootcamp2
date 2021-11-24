package main

import (
	"errors"
	"fmt"
)

func main() {
	const (
		minimo   = "minimo"
		promedio = "promedio"
		maximo   = "maximo"
	)
	notas1, err1 := operar(minimo, 2, 1, 4, 2, 7, 9)
	notas2, err2 := operar(maximo, 2, 1, 4, 2, 7, 9)
	notas3, err3 := operar(promedio, -2, 1, 4, 2, 7, 9)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(notas1)
	}
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(notas2)
	}
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(notas3)
	}

}

func operar(op string, notas ...int) (int, error) {
	/* Pasar aca una operacion y las notas, recuperar la funcion y usar el slice */
	resOperacion, err := operacion(op)
	if err != nil {
		return 0, err
	}
	res, err2 := resOperacion(notas...)
	if err2 != nil {
		return 0, err2
	}
	return res, nil

}

func operacion(op string) (func(notas ...int) (int, error), error) {
	switch op {
	case "minimo":
		return opMin, nil
	case "maximo":
		return opMax, nil
	case "promedio":
		return opPromedio, nil
	default:
		return nil, errors.New("operación incorrecta")
	}
}

func opPromedio(notas ...int) (int, error) {
	suma := 0
	for _, nota := range notas {
		suma += nota
		if nota < 0 {
			return 0, errors.New("no puede haber calificaciones negativas")
		}
	}
	return (suma / len(notas)), nil
}

func opMin(notas ...int) (int, error) {
	min := 10
	for _, nota := range notas {
		if nota < min {
			min = nota
		}
	}
	return min, nil
}

func opMax(notas ...int) (int, error) {
	max := 1
	for _, nota := range notas {
		if nota > max {
			max = nota
		}
	}
	return max, nil
}
