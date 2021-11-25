package main

import "fmt"

var word string = "summer"

func main() {
	fmt.Println("\nLa palabra tiene ", len(word), "letras")

	for _, letra := range word {
		fmt.Println(string(letra))
	}
}
