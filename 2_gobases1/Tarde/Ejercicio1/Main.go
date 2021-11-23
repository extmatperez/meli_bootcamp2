package main

import "fmt"

func main()  {
	
	var palabra string

	palabra = "PalabraLarga"

	fmt.Println(palabra)

	fmt.Println("La palabra tiene de largo", len(palabra), "letras")
	for _, letra := range palabra {
		fmt.Printf("%c  \n", letra)
	}
}