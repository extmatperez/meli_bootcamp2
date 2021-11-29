package main

import (
	"fmt"
	"os"
)

func readFile(fileName string) {

	content, err := os.ReadFile(fileName)

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado\n")
	}

	fmt.Println(string(content) + "\n")

}

func main() {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("ejecución finalizada")
	}()

	var fileName string
	fmt.Printf("Ingrese el nombre del archivo incluir el .txt: ")
	fmt.Scanf("%v", &fileName)
	fmt.Println("Nombre de archivo: ", fileName)

	readFile(fileName)
}
