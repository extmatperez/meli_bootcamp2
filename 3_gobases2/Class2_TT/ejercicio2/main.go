package main

import "fmt"

func main() {
	m := matrix{[]float64{}, 3.5, 4.5, true, 7.0}
	//values := {3.0,3.5,5.6,7.8,5.5,3.3,6.5,3.9,4.0}
	m.set(3.0, 3.5, 5.6, 7.8, 5.5, 3.3, 6.5, 3.9, 4.0)
	m.print(m)
}

type matrix struct {
	matriz     []float64
	alto       float64
	ancho      float64
	cuadratica bool
	maximo     float64
}

func (m *matrix) set(values ...float64) {
	m.matriz = values
}

func (m matrix) print(mat matrix) {
	fmt.Println(mat)
	fmt.Printf("%+v\n", mat)
}
