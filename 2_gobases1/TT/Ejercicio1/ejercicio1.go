package main

import (
	"fmt"
	"unicode"
)

func main() {
	una_variable := "Algunas letras"
	var las_letras []string
	for _, v := range una_variable {
		if !unicode.IsSpace(v) {
			las_letras = append(las_letras, string(v))
		}
	}
	fmt.Println(las_letras)
}
