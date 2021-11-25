/*
? Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de
productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.

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

	/*
		data, _ := os.ReadFile("./file.txt")

		if err == nil {
			file := string(data)
			fmt.Println(file)
		} else {
			fmt.Println("El archivo no existe...")
		} */

	p1 := Productos{001, 100.00, 10}
	p2 := Productos{002, 200.00, 20}
	p3 := Productos{003, 300.00, 30}
	p4 := Productos{004, 400.00, 40}

	//fmt.Printf("%d\n %10.2f\n %d\n", p1.Id, p1.Precio, p1.Cantidad)

	var lista []Productos

	lista = append(lista, p1)
	lista = append(lista, p2)
	lista = append(lista, p3)
	lista = append(lista, p4)

	p1for, _ := json.Marshal(lista)

	//textoFormateado := fmt.Sprintf("%v;", string(p1for))

	_ = os.WriteFile("./file.txt", p1for, 0644)

	//fmt.Println(textoFormateado)

	data, _ := os.ReadFile("./file.txt")

	var pListaLeida []Productos

	json.Unmarshal(data, &pListaLeida)

	fmt.Printf("\n%+v;", pListaLeida)
}
