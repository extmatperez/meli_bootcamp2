package main

import (
	"fmt"
)

type Matrix struct {
	rows [][]float64
}

func (m *Matrix) Set() {

	var row, col int
	var data float64

	fmt.Println("Ingrese el numero de filas: ")
	fmt.Scan(&row)
	fmt.Println("Ingrese el numero de columnas: ")
	fmt.Scan(&col)

	for i := 0; i < row; i++ {
		var localRow []float64
		for j := 0; j < col; j++ {
			fmt.Println("Ingrese el valor de la posicion: ", i, j)
			fmt.Scan(&data)
			localRow = append(localRow, data)
		}
		m.rows = append(m.rows, localRow)
	}
}

func (m Matrix) Print() {
	fmt.Println()
	for i := 0; i < len(m.rows); i++ {
		for j := 0; j < len(m.rows[i]); j++ {
			fmt.Print(m.rows[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
func main() {
	/*
		Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
		Para ello requieren una estructura Matrix que tenga los métodos:
		Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
		Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
		La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.
	*/
	var neo Matrix

	neo.Set()
	neo.Print()

}
