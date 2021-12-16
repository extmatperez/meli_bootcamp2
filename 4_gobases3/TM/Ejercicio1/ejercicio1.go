package main

import (
	"fmt"
)

func main() {
	var productos []Producto
	productos = append(productos, Producto{123, "Lavandina", 150, 25})
	productos = append(productos, Producto{124, "Detergente para manos", 110, 150})
	productos = append(productos, Producto{125, "Detergente para pisos", 250, 40})
	productos = append(productos, Producto{126, "Fregona", 350, 8})

	fmt.Println(productos)
}

type Producto struct {
	Id       int64   `json:"id"`
	Producto string  `json:"producto"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

func escribir_csv(headers string, data []string) {

}
