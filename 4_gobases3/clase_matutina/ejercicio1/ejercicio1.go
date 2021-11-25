package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const rutaArchivo = "/Users/aghione/Desktop/repositorios/bootcamp/practicas/meli_bootcamp2/4_gobases3/a.txt"

func main() {
	producto1 := Producto{1, 156.99, 5.0}
	producto2 := Producto{2, 171.00, 2.0}
	producto3 := Producto{3, 92.99, 1.0}
	producto4 := Producto{4, 15.50, 7.0}
	producto5 := Producto{5, 1563.00, 1.0}
	productos := []Producto{producto1, producto2, producto3, producto4, producto5}
	for _, producto := range productos {
		producto.Guardar()
	}
}

type Producto struct {
	Id       int64   `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad float64 `json:"cantidad"`
}

func (p Producto) Guardar() {
	prodFormateado, err := json.Marshal(p)
	textoFormateado := fmt.Sprintf("%v;", string(prodFormateado))

	archivo, err := os.OpenFile(rutaArchivo, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Algo salio mal al abrir el archivo.")
	}

	defer archivo.Close()

	if _, err = archivo.WriteString(textoFormateado); err != nil {
		fmt.Println("Algo salio mal al escribir el archivo.")
	} else {
		fmt.Println("Se guardo correctamente.")
	}
}

type Fichero interface {
	Guardar()
}
