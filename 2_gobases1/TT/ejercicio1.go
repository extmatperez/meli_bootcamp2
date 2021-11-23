package main

import "fmt"

func main() {
	palabra := "Palabra"
	fmt.Printf("%v tiene %v letras\n", palabra, len(palabra))

	for _, letra := range palabra {
		//letras = append(letras, string(letra))
		fmt.Println(string(letra))
	}
}