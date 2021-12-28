/*Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y
coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Products struct {
	ID       int     `json: "id"`
	Precio   float64 `json: "apellido"`
	Cantidad int     `json: "cantidad"`
}

func main() {
	var err error
	//err = os.Setenv("PRODUCTOS", "limpieza")
	//variable := os.Getenv("PRODUCTOS")
	//fmt.Println(variable)
	//fmt.Println(string(data))

	p1 := Products{1, 200, 3}
	p2 := Products{2, 500, 5}
	var lista []Products
	lista = append(lista, p1, p2)
	listaFormateado, err := json.Marshal(lista)

	err = os.WriteFile("./productos.txt", listaFormateado, 0644)
	data, err := os.ReadFile("./productos.txt")
	fmt.Println(err)
	//fmt.Println(string(data))

	var listRead []Products
	json.Unmarshal(data, &listRead)
	for i := 0; i < len(listRead); i++ {
		fmt.Printf("%+v,", listRead[i])
	}
	fmt.Println()

	//fmt.Printf("El archivo contiene:\n %+v\n", listRead[0])

}
