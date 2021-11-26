package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.ReadFile("./customers.txt")

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
			fmt.Println("Ejecución finalizada")
		}
	}()

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

	fmt.Println("Ejecución finalizada")
}
