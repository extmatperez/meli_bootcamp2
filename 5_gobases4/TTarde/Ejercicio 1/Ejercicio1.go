package main

import (
	"fmt"
	"os"
)

func leerArchivo(nombre string){
	data, err := os.ReadFile(nombre)
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

	fmt.Println(data)

}

func main(){
leerArchivo("customers.txt")
fmt.Println("Ejecución completada")
}