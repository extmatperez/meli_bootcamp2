/*
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados,
separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
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

	producto1 := Producto{111223, 30012.00, 1}
	producto2 := Producto{444321, 1000000.00, 4}
	producto3 := Producto{434321, 50.50, 1}

	var lista []Producto

	lista = append(lista, producto1)
	lista = append(lista, producto2)
	lista = append(lista, producto3)

	listaFormateada, err := json.Marshal(lista)

	err = os.WriteFile("../archivos/myFile.txt", listaFormateada, 0644)

	if err != nil {
		fmt.Println("No se pudo escribir")
	}

}
