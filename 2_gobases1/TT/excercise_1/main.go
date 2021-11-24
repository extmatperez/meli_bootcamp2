package main

import "fmt"

func main() {

	var palabra string
	fmt.Println("ingrese la palabra:")
	fmt.Scanln(&palabra)

	fmt.Println("Esta palabra tiene ", len(palabra), "caracter(es)")
	fmt.Println("La palabra deletreada es:")
	for _, texto := range palabra {
		fmt.Println(string(texto))
	}
}
