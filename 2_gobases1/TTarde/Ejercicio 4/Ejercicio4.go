package main

import "fmt"

func main() {

	/*Forma 1: Switch. Codigo mas largo y menos amigable,
	pero menos uso de memoria (ya que solo tenemos una variable ocupando memoria [mes]).
	*/
	var mes int

	fmt.Println("Ingrese numero de mes")
	fmt.Scanln(&mes)

	switch mes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("Numero de mes no válido")

	}

	/*Forma 2: array. Codigo más corto y amigable, pero ocupa más memoria ya que necesita
	cargar también un array de strings.
	*/

	mes_v := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	fmt.Println("Ingrese numero de mes")
	fmt.Scanln(&mes)

	if mes > 0 && mes < 11 {
		fmt.Println(mes_v[mes-1])
	} else {
		fmt.Println("Numero de mes no válido")
	}
}
