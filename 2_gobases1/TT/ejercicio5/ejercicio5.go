/*Una profesora de la universidad quiere tener un listado con todos sus estudiantes.
Es necesario crear una aplicación que contenga dicha lista.
Estudiantes:
Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernán, Leandro, Eduardo, Duvraschka.

Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado,
sin modificar el código que escribiste inicialmente.
Estudiante:
Gabriela*/

package main

import "fmt"

func main() {

	slice := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Println("Listado Original: ")
	for _, elem := range slice {
		fmt.Println(elem)
	}
	fmt.Println("")
	fmt.Println("Listado Modificado: ")
	slice = append(slice, "Gabriela")

	for _, elem := range slice {
		fmt.Println(elem)
	}
}
