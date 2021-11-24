package main

import "fmt"

type Matrix struct {
	values      [][]float64
	maxvalue    float64
	rows        int
	columns     int
	isCuadratic bool
}

func (m *Matrix) setValues(rows int, columns int, values ...float64) {
	m.rows = rows
	m.columns = columns
	max := 0.0
	//setea valores en la matriz
	//por cada valor pregunto si es mayor a max, si es, lo guardo en maxValue
	m.maxvalue = max
	m.isCuadratic = len(m.values) == len(m.values[0]) //es o no cuadrada
}

func (m Matrix) printValues() {
	for i, value := range m.values {
		fmt.Println("Row #", i, ": ", value)
	}
}

func main() {
	newMatrix := Matrix{}
	newMatrix.setValues(2, 2, 2.0, 3.2, 1.7, 4.0)
	newMatrix.printValues()
}
