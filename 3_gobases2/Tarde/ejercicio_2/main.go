package main

import (
	"fmt"
)

type Matrix struct {
	alto       int
	ancho      int
	cuadratica bool
	maximo     float64
	contenido  [3][3]float64
}

func (m *Matrix) set(x int, y int, cu bool) {
	m.alto = y
	m.ancho = x
	m.cuadratica = cu
	max := 0.0
	var intArr [3][3]float64
	fmt.Println("Ingrese los valores para la matriz 3x3: ")

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			num := 0.0
			fmt.Scanf("%f", &num)
			intArr[i][j] = num
			if num > max {
				max = num
			}
		}
		m.maximo = max
		m.contenido = intArr

	}

}

func (m Matrix) print() {
	fmt.Println("Print: ")
	for i := 0; i < len(m.contenido); i++ {
		for j := 0; j < len(m.contenido[i]); j++ {
			fmt.Print(m.contenido[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Printf("maximo= %v\n", m.maximo)
}

func main() {
	m1 := Matrix{}
	m1.set(3, 3, false)
	m1.print()

}
