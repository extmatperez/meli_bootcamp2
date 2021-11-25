/*Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una
estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la
estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de
	línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto,
la dimensión del ancho, si es cuadrática y cuál es el valor máximo.*/

package main

import "fmt"

type Matrix struct {
	alto       int
	ancho      int
	valores    []float64
	cuadrática bool
	maximo     float64
}

func (m *Matrix) Set(alto, ancho int, cuadrática bool, valores ...float64) {
	m.alto = alto
	m.ancho = ancho
	m.valores = valores
	m.cuadrática = cuadrática
}

func (m *Matrix) Print() {

	index := 0

	for i := 0; i < len(m.valores); i++ {

		if index == m.ancho {
			fmt.Printf("\n%f ", m.valores[i])
			index = 0
		} else {
			fmt.Printf("%f ", m.valores[i])
			index++
		}
	}
	fmt.Println("")
}

func main() {

	m1 := Matrix{}

	m1.Set(2, 2, true, 3.4, 4, 7.8, 1.2)

	m1.Print()
}
