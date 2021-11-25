package main

import (
	"fmt"
	"os"
)

type Producto struct {
	id       int
	precio   float64
	cantidad int
}

func main() {
	var productos []Producto

	producto1 := Producto{1, 1000.0, 2}
	producto2 := Producto{2, 2000.0, 3}
	producto3 := Producto{3, 3000.0, 4}
	productos = append(productos, producto1, producto2, producto3)
	var datos string
	for _, p := range productos {
		datos += fmt.Sprintf("%d;%.2f;%d \n", p.id, p.precio, p.cantidad)
	}
	d1 := []byte(datos)

	err := os.WriteFile("./archivos/datos.csv", d1, 0644)

	if err == nil {
		fmt.Println("Guardado correctamente")
	} else {
		fmt.Println("Ocurrio un problema")
	}

}
