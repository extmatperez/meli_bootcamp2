package main

import "fmt"

var palabra string = "frambuesa"
var deletreo []string

func main() {
	for _, letra := range palabra {
		deletreo = append(deletreo, string(letra))
	}
	fmt.Println("Cantidad de letras ", len(deletreo))
	fmt.Println(deletreo)
}
