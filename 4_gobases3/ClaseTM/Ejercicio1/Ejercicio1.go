package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Producto struct {
	Id       int
	Precio   float64
	Cantidad int
}

func main() {
	p1 := Producto{3, 100.50, 1}

	var csvProducts string = "ID:1,PRECIO:20,CANTIDAD:4 \n ID:2,PRECIO:10,CANTIDAD:4"
	guardarArchivoCsv(csvProducts)

	guardarArchivo(p1)
}

func guardarArchivoCsv(csv string) {

	miMarsh, err := json.Marshal(csv)

	if err == nil {
		err := os.WriteFile("./ProductosCSVHardcoded.txt", miMarsh, 0644)

		if err == nil {
			fmt.Println("Se agrego correctamente")
		}
	}

}

func guardarArchivo(producto Producto) {

	data, err := os.ReadFile("./Productos.txt")

	if err == nil {
		// productos = append(productos, string(data))
		var historialProd []Producto

		err := json.Unmarshal(data, &historialProd)
		fmt.Println(err)

		fmt.Println(historialProd)

		historialProd = append(historialProd, producto)

		miMarshHistorial, err := json.Marshal(historialProd)

		if err == nil {
			err := os.WriteFile("./Productos.txt", miMarshHistorial, 0644)
			fmt.Println(err)
		}

	} else {
		miMarsh, err := json.Marshal(producto)

		if err == nil {
			err := os.WriteFile("./Productos.txt", miMarsh, 0644)
			fmt.Println(err)

		}

	}

}
