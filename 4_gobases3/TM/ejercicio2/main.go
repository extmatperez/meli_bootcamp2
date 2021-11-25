/*
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por
pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y
a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe
visualizar el total (Sumando precio por cantidad)
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Producto struct {
	Id       int     `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

func main() {

	data, err := os.ReadFile("../archivos/myFile.txt")

	if err != nil {
		fmt.Println("No se pudo leer")
	} else {
		var listaLeida []Producto

		json.Unmarshal(data, &listaLeida)

		var total float64

		fmt.Printf("%-12v%12v%12v\n", "ID", "Precio", "Cantidad")

		for _, p := range listaLeida {
			fmt.Printf("%-12v%12.2f%12v\n", p.Id, p.Precio, p.Cantidad)

			total += p.Precio
		}

		fmt.Printf("%24.2f\n", total)
	}
}
