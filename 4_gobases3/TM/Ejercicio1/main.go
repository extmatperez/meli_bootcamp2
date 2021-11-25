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
	p1formateado, err := json.Marshal(p1)

	err = os.WriteFile("./productos.txt", p1formateado, 0644)
	data, err := os.ReadFile("./productos.txt")
	fmt.Println(err)
	fmt.Println(string(data))

}
