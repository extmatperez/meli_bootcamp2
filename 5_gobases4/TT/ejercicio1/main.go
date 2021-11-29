package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("ARCHIVO", "customers.txt")

	LeerArchivo()
	fmt.Printf("Ejecución finalizada\n")

}

func LeerArchivo() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.ReadFile("./" + os.Getenv("ARCHIVO"))

	if err != nil {
		panic("Error reading file " + os.Getenv("ARCHIVO"))
	}
	fmt.Printf("File read.\n")
}
