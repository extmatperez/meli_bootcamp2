package main

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	x int
	y int
	m [][]float64
}

func (m *Matrix) makeMatriz() {
	a := make([][]float64, m.x)
	for i := range a {
		a[i] = make([]float64, m.y)
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = (0 + rand.Float64()*(100-0))
		}
	}
	m.m = a
}

func main() {
	mTest := Matrix{x: 3, y: 5}
	mTest.makeMatriz()

	for i := 0; i < len(mTest.m); i++ {
		fmt.Println(mTest.m[i])
	}

}
