/*
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de
calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo
y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar
(mínimo, máximo o promedio) y que devuelva otra función (y un error en caso que el cálculo
no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo
que se indicó en la función anterior
*/

package main

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)


func funMinimo(calificaciones ...float64) float64 {
	primero := true
	min := 0.0

	for _, c := range calificaciones{
		if primero {
			primero = false
			min = c
		} else if c < min {
			min = c
		}
	}

	return min
}

func funPromedio(calificaciones ...float64) float64{
	acumulador := 0.0

	for _, c := range calificaciones{
		acumulador += float64(c)
	}

	return acumulador / float64(len(calificaciones))
}

func funMaximo(calificaciones ...float64) float64{
	primero := true
	max := 0.0

	for _, c := range calificaciones{
		if primero {
			primero = false
			max = c
		} else if c > max{
			max = c
		}
	}

	return max
}


func operacion(tipoOperacion string) (func(calificaciones ...float64) float64, error){

	if tipoOperacion != "minimo" && tipoOperacion != "promedio" && tipoOperacion != "maximo" {
		return nil, errors.New("cálculo no definido")
	}

	switch tipoOperacion {
	case "minimo":
		return funMinimo, nil
	case "promedio":
		return funPromedio, nil
	case "maximo":
		return funMaximo, nil
	}

	return funMinimo, nil
}


func main() {
	
	minFunc, err1 := operacion(minimo)
	promFunc, err2 := operacion(promedio)
	maxFunc, err3 := operacion(maximo)

	if err1 != nil{
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("Valor mínimo: %v\n", minFunc(2,3,3,4,1,2,4,5))
	}

	if err2 != nil{
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("Valor promedio: %v\n", promFunc(2,3,3,4,1,2,4,5))
	}

	if err3 != nil{
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("Valor máximo: %v\n", maxFunc(2,3,3,4,1,2,4,5))
	}
}