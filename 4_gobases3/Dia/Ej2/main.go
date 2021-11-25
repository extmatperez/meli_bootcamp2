/*
? Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50

*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Productos struct {
	Id       int     `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

func main() {

	data, _ := os.ReadFile("../Ej1/file.txt")

	var pListaLeida []Productos

	json.Unmarshal(data, &pListaLeida)

	fmt.Printf("%-10v %15v %10v\n", "Id", "Precio", "Cantidad")

	sum := 0.0
	for i := 0; i < len(pListaLeida); i++ {

		fmt.Printf("%-10v %15v %10v\n", pListaLeida[i].Id, pListaLeida[i].Precio, pListaLeida[i].Cantidad)

		sum = sum + pListaLeida[i].Precio
	}

	fmt.Printf("%26v\n ", sum)

}
