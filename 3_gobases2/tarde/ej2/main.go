package main

import (
	"fmt"
	"math"
)

func roundToInt(x float64) int {
	return int(math.Round(x))
}

type Matrix struct {
	values   []float64
	height   int
	width    int
	isSquare bool
	maxValue float64
}

func (mat *Matrix) Set(values ...float64) {
	var count float64 = float64(len(values))

	var square = math.Sqrt(count)
	var height int = roundToInt(square)

	var div = count / float64(height)
	var width int = roundToInt(div)

	if float64(width*height) < count {
		width++
	}

	var maxValue float64 = -999999999.0
	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	mat.values = values
	mat.height = height
	mat.width = width
	mat.isSquare = (height == width) && (count == float64(height*width))
	mat.maxValue = maxValue
}

func (mat Matrix) Print() {
	var init, end int = 0, mat.width

	for i := 1; i < mat.height; i++ {
		fmt.Printf("Fila %d: %v\n", i, mat.values[init:end])
		init = end
		end = end + mat.width
	}

	end = len(mat.values)
	fmt.Printf("Fila %d: %v\n", mat.height, mat.values[init:end])
}

func main() {
	var mat Matrix
	mat.Set(4, 5, 8, 654, 8, 123, 1, 5, 8, 46, 6, 8, 4563, 6, 84, 67)
	mat.Print()
	fmt.Printf("%+v\n", mat)
}
