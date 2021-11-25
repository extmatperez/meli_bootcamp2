// Ejercicio 1 - Guardar archivo
// Una empresa que se encarga de vender productos de limpieza necesita:
// Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
// Debe tener el id del producto, precio y la cantidad.
// Estos valores pueden ser hardcodeados o escritos en duro en una variable.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// func generar_informe_productos_comprados() {

// }

func main() {

	type Producto struct {
		Id       int `json:"id"`
		Precio   int `json:"precio"`
		Cantidad int `json:"cantidad"`
	}
	//productos := []string{"1.2000.1", "2.1000.3", "3.222.1", "4.3000.4"}

	compras1 := Producto{1, 2000, 2}
	compras2 := Producto{2, 3000, 2}
	compras3 := Producto{4, 2256, 4}

	var lista_compras []Producto

	lista_compras = append(lista_compras, compras1)
	lista_compras = append(lista_compras, compras2)
	lista_compras = append(lista_compras, compras3)
	//Esto devuelve un slice de bytes y un error
	compras1_formart, err := json.Marshal(lista_compras)

	//productos_comprados2 := Producto{2, 3000, 1}
	//productos_comprados3 := Producto{3, 5000, 2}
	//fmt.Println(productos_comprados1)
	//producto1 := []byte()
	// productos_tobyte := []byte(productos)
	err = os.WriteFile("./compras.txt", compras1_formart, 0644)
	if err != nil {
		log.Fatal(err)
		fmt.Println("No")
	}

}
