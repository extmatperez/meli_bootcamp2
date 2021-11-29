package main

import (
	"fmt"
)

// Ejercicio 2 - Matrix
// Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
// Para ello requieren una estructura Matrix que tenga los métodos:
// Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
// Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
// La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.

type Matrix struct {
	Values    []float64
	Height    int
	Width     int
	Quadratic bool
	MaxValue  int
}

func (m *Matrix) set(values []float64) {
	m.Values = values
}

func (m *Matrix) printMatrix() {

	if m.Values != nil {
		i := 0
		counter := 0
		for i < len(m.Values) {

			fmt.Printf("%v ", m.Values[i])
			i++
			counter++
			if counter == m.Width {
				fmt.Printf("\n")
				counter = 0
				continue
			}
		}

	}

}

func main() {

	// var matriz1 Matrix
	// matriz1.Height = 2
	// matriz1.Width = 3

	// var lista = []float64{1, 2, 3, 4, 5, 6}
	// matriz1.set(lista)

	// fmt.Printf("matriz nueva: %v\n", matriz1.Values)

	// fmt.Print("matriz 1 impresa: \n")
	// matriz1.printMatrix()

	var matriz2 Matrix
	matriz2.Width = 3

	var lista2 = []float64{10, 20, 30, 40, 50, 60, 70, 80, 90}
	matriz2.set(lista2)

	fmt.Print("matriz 2 impresa: \n")
	matriz2.printMatrix()
}
