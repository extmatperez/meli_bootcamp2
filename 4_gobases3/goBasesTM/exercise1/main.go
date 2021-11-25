package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	producto := Producto{1122, 300.10, 5}
	producto1 := Producto{1123, 303.10, 10}
	producto2 := Producto{1124, 333.10, 15}

	/* 	var lista []Producto
	   	lista = append(lista, producto, producto1, producto2) */

	guardarArchivo(producto, producto1, producto2)

	fmt.Println("\n", producto, "\n", producto1, "\n", producto2)

}

type Producto struct {
	IdProd   int     `json:"ID"`
	Precio   float64 `json:"PRECIO"`
	Cantidad int     `json:"CANTIDAD"`
}

func guardarArchivo(prod ...Producto) {

	formProd, err := json.Marshal(prod)

	if err != nil {
		fmt.Println("Error json: ", err)
	}

	prodFile := os.WriteFile("./fileE.txt", formProd, 0644)

	if prodFile != nil {
		fmt.Println("Error", err)
	}

	//chequear esto de abajo

	data, err := os.ReadFile("./fileE.txt")
	if err == nil {
		file := string(data)
		fmt.Println(file)
	} else {
		fmt.Println("El archivo no existe...")
	}

}
