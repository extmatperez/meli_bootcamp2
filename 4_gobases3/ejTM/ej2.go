package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Producto struct {
	Id       string  `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

func leer_productos_de_archivo(ruta_entrada string) []Producto {
	archivo_csv, errorr := os.Open(ruta_entrada)
	if errorr != nil {
		fmt.Println("error al abrir el archivo: ", errorr)

	}
	defer archivo_csv.Close()
	lector_csv := csv.NewReader(archivo_csv)
	lector_csv.Comma = ';'

	lista_valores_csv, err := lector_csv.ReadAll()
	if err != nil {
		fmt.Println("error al abrir el archivo: ", err)
	}
	var lista_productos = []Producto{}
	for _, product := range lista_valores_csv {
		id, precio, cantidad := product[0], product[1], product[2]
		float_precio, _ := strconv.ParseFloat(precio, 64)
		int_cantidad, _ := strconv.ParseInt(cantidad, 10, 64)
		lista_productos = append(lista_productos, Producto{id, float_precio, int(int_cantidad)})
	}

	return lista_productos
}

func main() {

	lista_de_productos := leer_productos_de_archivo("productos.txt")
	productos, _ := json.Marshal(lista_de_productos)
	fmt.Println(string(productos))

	fmt.Println("valor:", lista_de_productos[0].Cantidad)
}
