package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Producto struct {
	id string
	price float64
	quantity int
}

func main() {
	// Leemos los productos
	productos := leerProductos()
	fmt.Printf("%-20s%-20s%-20s\n", "ID", "Precio", "Cantidad")
	totalPrice := 0.0
	for _, p := range productos {
		// Imprimimos el producto
		fmt.Printf("%-20s%-20.2f%-20d\n", p.id, p.price, p.quantity)
		totalPrice += p.price * float64(p.quantity)
	}
	fmt.Printf("%-20s%-20.2f%-20s\n", "", totalPrice, "")
}

func leerProductos() []Producto {
	f, err := os.Open("products.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var productos []Producto
	for _, row := range rows {
		parsedPrice, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			panic(err)
		}
		parsedQuantity, err := strconv.Atoi(row[2])
		if err != nil {
			panic(err)
		}
		p := Producto{
			id: row[0],
			price: parsedPrice,
			quantity: parsedQuantity,
		}
		productos = append(productos, p)
	}
	return productos
}