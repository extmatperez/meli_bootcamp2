/* La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores
tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y
abajo del precio se debe visualizar el total (Sumando precio por cantidad)
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
	var cantTotal int
	var precioTotal float64
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

	fmt.Printf("%-10v", "ID")
	fmt.Printf("%10v", "Precio")
	fmt.Printf("%11v", "Cantidad\n")

	for _, value := range listRead {

		fmt.Printf("%-10v", value.ID)
		fmt.Printf("%10v", value.Precio)
		fmt.Printf("%10v", value.Cantidad)
		fmt.Println()
		cantTotal += value.Cantidad
		precioTotal += value.Precio

	}
	fmt.Printf("%20v", precioTotal)
	fmt.Printf("%10v", cantTotal)
	/* for i := 0; i < len(listRead); i++ {
		fmt.Printf("%+v,", listRead[i])
	} */
	fmt.Println()

	//fmt.Printf("El archivo contiene:\n %+v\n", listRead[0])

}
