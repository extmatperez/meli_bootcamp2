package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Productos struct {
	Id       int     `json:"id"`
	Precios  float64 `json:"precios"`
	Cantidad int     `json:"cantidad"`
}

func main() {
	/*
		La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)
		Ejemplo:
		ID 		Precio 		Cantidad
		111223 	30012.00 	1
		444321 	1000000.00 	4
		434321 	50.50 		1
				4030062.50
	*/

	var dataF []Productos
	data, _ := os.ReadFile("../archivo/myFile1.txt")
	err := json.Unmarshal(data, &dataF)

	if err != nil {
		fmt.Println(err)
	} else {
		var total float64
		fmt.Printf("\n%-20s%15s%12s\n", "ID", "Precio", "Cantidad")
		for _, value := range dataF {
			fmt.Printf("%-20d%15.2f%12d\n", value.Id, value.Precios, value.Cantidad)
			total += value.Precios * float64(value.Cantidad)
		}

		fmt.Println(total)
	}
}
