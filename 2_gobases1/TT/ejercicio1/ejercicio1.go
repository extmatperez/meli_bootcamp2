package main

import "fmt"

func main() {
	palabra := "Palabra"
	fmt.Printf("%v tiene %v letras\n", palabra, len(palabra))

	for _, letra := range palabra {
		fmt.Println(string(letra))
	}
}