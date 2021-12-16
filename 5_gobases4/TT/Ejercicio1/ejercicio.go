package main

import (
	"fmt"
	"os"
)

func main() {
	file := leer_archivo("./customers.txt")
	fmt.Println(file)
}

func leer_archivo(file_path string) os.File {
	file, err := os.OpenFile(file_path, os.O_RDWR, 0644)
	defer func() {
		fmt.Println("Ejecucion finalizada")
	}()
	if err != nil {
		panic("El archivo indicado no fue encontrado o esta dañado\n")
	}
	return *file
}
