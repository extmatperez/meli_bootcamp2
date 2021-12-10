package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("Ejecucion terminada")
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	data, err := os.ReadFile("./customers.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	fmt.Println(data)
}
