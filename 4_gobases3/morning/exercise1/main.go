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
		Una empresa que se encarga de vender productos de limpieza necesita:
		Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
		Debe tener el id del producto, precio y la cantidad.
		Estos valores pueden ser hardcodeados o escritos en duro en una variable.
	*/
	//productos := []productos{{1, 10.50, 4}, {2, 99.40, 8}, {3, 5.50, 2}}
	producto1 := Productos{1, 10.50, 4}
	producto2 := Productos{2, 99.40, 8}
	producto3 := Productos{3, 5.50, 2}

	var produc []Productos

	produc = append(produc, producto1)
	produc = append(produc, producto2)
	produc = append(produc, producto3)

	pFormateado, err := json.Marshal(produc)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Print("\n", string(pFormateado), "\n", "\n")
	}
	err = os.WriteFile("../archivo/myFile1.txt", pFormateado, 0644)
}
