package main

import (
	"fmt"
	"os"
)

type Producto struct {
	Id       int     `json:id`
	Precio   float64 `json:precio`
	Cantidad int     `json:cantidad`
}

func main() {
	prod1 := Producto{123, 100.00, 2}
	prod2 := Producto{111, 1771.20, 4}
	prod3 := Producto{222, 143.22, 1}
	stringCompleto := escribirProductos(prod1, prod2, prod3)
	err := os.WriteFile("./productos.csv", []byte(stringCompleto), 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El archivo se ha escrito con exito")
	}
}

func escribirProductos(productos ...Producto) string {
	values := ""
	for _, p := range productos {
		values += fmt.Sprintf("%v;%.2f;%v\n", p.Id, p.Precio, p.Cantidad)
	}
	return values
}
