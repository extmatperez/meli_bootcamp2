package main

import "fmt"

func main() {
	var palabra = "Cualquiera"
	fmt.Println("Palabra:", palabra)
	fmt.Printf("La palabra tiene %d carateres", len(palabra))
	fmt.Printf("\nDeletreada: ")
	for _,c := range palabra {
		fmt.Printf("%c ", c)
	}
}