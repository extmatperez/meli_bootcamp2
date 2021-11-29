package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("custom.txt")

	defer func() {
		fmt.Println("Ejecución finalizada")
	}()

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

}
