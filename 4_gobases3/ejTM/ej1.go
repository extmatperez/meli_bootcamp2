package main

import (
	"fmt"
	"os"
)

// type Producto struct {
// 	Id       string  `json:"id"`
// 	Precio   float64 `json:"precio"`
// 	Cantidad int     `json:"cantidad"`
// }

func escribir_productos_en_archivo() {

	ruta := "./productos.txt"

	producto1 := Producto{"1234543", 34.4, 5}
	producto2 := Producto{"1277878", 123300.0, 3}
	producto3 := Producto{"8765435", 123.0, 1}

	var slice_productos = []Producto{producto1, producto2, producto3}
	os.WriteFile(ruta, []byte(""), 0644)

	for _, producto := range slice_productos {
		linea := fmt.Sprintf("%s;%.1f;%d\n", producto.Id, producto.Precio, producto.Cantidad)
		file, err := os.OpenFile(ruta, os.O_WRONLY|os.O_APPEND, 0644)
		if err == nil {
			file.Write([]byte(linea))
		}
		file.Close()
	}

}
