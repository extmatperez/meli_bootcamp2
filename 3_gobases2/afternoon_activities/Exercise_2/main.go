/* Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

package main

import "fmt"

// Método Set: recibe valores float64 e inicializa valores
// Método Print: imprime la matriz con saltos de línea entre filas

type Matrix struct {
	Values    []float64
	High      float64
	Width     float64
	Quadratic bool
	MaxValue  float64
}

func (m *Matrix) Set(high, width float64, values ...float64) {
	m.High = high
	m.Width = width

	if m.High == m.Width {
		m.Quadratic = true
	}
	m.MaxValue = m.High * m.Width
	m.Values = values
}
func (m *Matrix) Print() {
	var high float64
	var width float64
	for _, val := range m.Values {
		if width == m.Width {
			fmt.Printf("\n")
			width = 0
			high++
		}
		if high == m.High {
			break
		}
		fmt.Printf("%.1f\t", val)
		width++
	}
}

func main() {
	var matrix Matrix
	matrix.Set(2, 3, 5, 7, 12, 78, 34, 4, 45, 45)
	matrix.Print()
}
