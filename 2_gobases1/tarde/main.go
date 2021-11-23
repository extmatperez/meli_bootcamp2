package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//Ejercicio 1
	reader := bufio.NewReader(os.Stdin)
	var palabra string
	palabra, _ = reader.ReadString('\n')
	fmt.Printf("%v%T\n", palabra, palabra)
	var letras []string = strings.Split(palabra, "")
	fmt.Printf("%v", letras)
	fmt.Printf("La palabra tiene %d letras\n", len(letras))
	fmt.Print("Palabra deletreada: ")
	for _, letra := range letras {
		fmt.Printf("%s ", letra)
	}
}
