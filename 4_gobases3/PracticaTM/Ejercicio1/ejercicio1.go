package main

import (
	"fmt"
	"os"
)

type Producto struct {
	Id       int     `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

func main() {
	var productos []Producto
	productos = append(productos, Producto{1, 1000.50, 10})
	productos = append(productos, Producto{2, 3050000.0, 5})
	productos = append(productos, Producto{3, 20.6, 2})
	productos = append(productos, Producto{4, 33.5, 10})
	productos = append(productos, Producto{5, 28.4, 20})

	var prodActual, stringProductos string
	stringProductos = "ID;Precio;Cantidad\n"
	for _, producto := range productos {
		prodActual = fmt.Sprintf("%v;%v;%v\n", producto.Id, producto.Precio, producto.Cantidad)
		stringProductos += prodActual
	}
	os.WriteFile("./ProductosComprados.csv", []byte(stringProductos), 0644)

	// prodFormateados, err := json.Marshal(productos)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	os.WriteFile("./ProductosComprados.csv", prodFormateados, 0644)
	// }

}
