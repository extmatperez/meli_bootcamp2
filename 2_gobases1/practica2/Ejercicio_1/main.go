package main

import "fmt"

func main(){

	// Declare variables
	var (
		word string = "Palabra"
	)
	fmt.Println("La cantidad de letras de la palabra son: ", len(word))
	fmt.Println("Las letras son: ")
	for i := 0; i < len(word); i++{
		fmt.Println(word[i:i+1])
	}
}