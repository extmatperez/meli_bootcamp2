package main

import "fmt"

type Matrix struct {
	mat        []float64
	alto       float64
	ancho      float64
	max        float64
	cuadratica bool
}

func (m *Matrix) set(values ...float64) {
	m.mat = values
}

func (m Matrix) print(mat Matrix) {
	fmt.Println(mat)
	fmt.Printf("%v\n", mat)
}

func main() {
	m := Matrix{[]float64{}, 2, 3.2, 4.7, true}
	m.set(2, 7.2, 3, 7, 8.5, 4.3, 9, 3)
	m.print(m)
}
