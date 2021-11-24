package main

import "fmt"

/*
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente
 una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho,
si es cuadrática y cuál es el valor máximo.
*/

type Matrix struct {
	valores    []float64
	alto       int
	ancho      int
	cuadratica bool
	max        int
}

func (m *Matrix) Set(valores ...float64) {
	m.valores = valores
	//falta setear el max acá
}

func (m Matrix) Print() {
	for i := 0; i < m.alto; i++{
		for j := 0; j < m.ancho, j++
			
	}
}

func main() {

	matriz1 := Matrix{alto: 3,
		ancho:      3,
		cuadratica: true}
	matriz1.Set(1, 2, 3, 4)
	fmt.Printf("%v", matriz1)
}
