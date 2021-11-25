/*
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear
una estructura que represente una matriz de datos.

Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la
estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de
línea entre filas)

La estructura Matrix debe contener los valores de la matriz, la dimensión del alto,
la dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

package main

import "fmt"

type Matrix struct {
	valores     [][]float64
	alto        int
	ancho       int
	cuadrática  bool
	valorMaximo float64
}

func (m *Matrix) SetMatrix(valores [][]float64) {
	
	m.valores = valores

	m.alto = len(valores)
	m.ancho = len(valores[0])

	if m.alto == m.ancho {
		m.cuadrática = true
	} else {
		m.cuadrática = false
	}

	m.valorMaximo = 0

	for i := 0; i < m.alto; i++ {
		for j := 0; j < m.ancho; j++ {
			if m.valorMaximo < m.valores[i][j] {
				m.valorMaximo = m.valores[i][j]
			}
		}
	}
}

func (m Matrix) PrintMatrix() {
	
	for i := 0; i < m.alto; i++ {

		for j := 0; j < m.ancho; j++ {
			fmt.Printf("%.f ", m.valores[i][j])
		}

		fmt.Printf("\n")
	}
}

func main() {

	var m Matrix

	m.SetMatrix([][]float64{	{1, 2},
								{3, 4},
								{5, 6},
							})

	m.PrintMatrix()
}
