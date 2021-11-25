package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Producto struct {
	ID       int
	Precio   float64
	Cantidad int
}

func main() {
	os.Setenv("ARCHIVO", "./productos.csv")
	escribirArchivo()
	leerArchivo()

}

func escribirArchivo() {
	// var productos []Producto
	// productos = append(productos, Producto{111223, 30012.00, 1})
	// productos = append(productos, Producto{444321, 1000000.00, 4})
	// productos = append(productos, Producto{444327, 50.50, 1})

	// for i := 0; i < 20; i++ {
	// 	productos = append(productos, generarProductoRandom())
	// }

	productos := crearProductos(3)
	productos = append(productos, crearProductos(5)...)

	f, err := os.OpenFile(os.Getenv("ARCHIVO"), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err == nil {
		f.Write([]byte("ID;Precio;Cantidad\n"))

		for _, producto := range productos {
			s := fmt.Sprintf("%d;%.2f;%d\n", producto.ID, producto.Precio, producto.Cantidad)
			f.Write([]byte(s))
		}
	}

	f.Close()
}

func leerArchivo() {
	f, _ := os.Open(os.Getenv("ARCHIVO"))
	reader := csv.NewReader(f)
	reader.Comma = ';'
	csvLines, _ := reader.ReadAll()
	fmt.Println(len(csvLines))
	precioTotal := 0.0
	for i, fila := range csvLines {
		if i == 0 {
			fmt.Printf("%-10s\t%10s\t%10s", fila[0], fila[1], fila[2])
		} else {
			id, _ := strconv.Atoi(fila[0])
			precio, _ := strconv.ParseFloat(fila[1], 64)
			cantidad, _ := strconv.Atoi(fila[2])
			fmt.Printf("%-10d\t%10.2f\t%10d", id, precio, cantidad)
			precioTotal += precio * float64(cantidad)
		}
		fmt.Println()
	}
	//fmt.Printf("\nMonto total:\t%10.2f\t\n\n", precioTotal)
	fmt.Printf("\n\t\t%10.2f\t\n\n", precioTotal)
}

func crearProductos(cantidad int) []Producto {
	var productos []Producto
	for i := 0; i < cantidad; i++ {
		productos = append(productos, generarProductoRandom())
	}
	return productos
}

func generarProductoRandom() Producto {
	ID := rand.Intn(10000)
	precio := rand.Float64() * 1000000
	cantidad := rand.Intn(10)
	return Producto{ID, precio, cantidad}
}
