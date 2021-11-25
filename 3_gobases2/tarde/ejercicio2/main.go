package main

import (
	"fmt"
)

func main() {
	//threedim := make([][][]float64, 1*1*1)
	threedim := [][][]float64{
		{{1, 2, 3}, {4, 5, 6}},
		{{7, 8, 9}, {10, 11, 12}},
		{{13, 14, 15}, {16, 17, 18}},
	}

	setData(threedim)
	printMatrix(threedim)

}

func setData(a [][][]float64) {
	a[0][0][0] = 0
	a[0][0][1] = 0
	a[0][0][2] = 0
}

func printMatrix(a [][][]float64) {

	fmt.Println("alto :", len(a), " ancho :", len(a[0]), " profundo :", len(a[0][0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			for k := 0; k < len(a[0][0]); k++ {
				fmt.Println(a[i][j][k])

			}
		}
	}
}
