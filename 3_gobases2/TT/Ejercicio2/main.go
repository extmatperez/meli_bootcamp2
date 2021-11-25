package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Matrix struct {
	Valores []float64
	Alto    int
	Ancho   int
	//Cuadrada bool
	//Max float64
}

func (m *Matrix) setMatrix(valores []float64) {
	m.Valores = valores
}

func (m Matrix) printMatrix() {
	cadena := ""
	for i, value := range m.Valores {
		cadena += fmt.Sprintf("%.2f  ", value)
		if (i+1)%m.Ancho == 0 {
			cadena += "\n"
		}
	}
	fmt.Println(cadena)
}

func main() {
	var valores []float64
	var alto int = 3
	var ancho int = 5

	myMatrix := Matrix{valores, alto, ancho}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < alto*ancho; i++ {
		valores = append(valores, math.Floor(r1.Float64()*100)/100)
	}

	myMatrix.setMatrix(valores)
	fmt.Println(myMatrix)

	myMatrix.printMatrix()

}
