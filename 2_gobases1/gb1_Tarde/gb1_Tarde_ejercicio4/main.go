package main

import "fmt"

// Ejercicio 4 - A qué mes corresponde
// Realizar una aplicación que contenga una variable con el número del mes. 
// Según el número, imprimir el mes que corresponda en texto. 
// ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
// Ej: 7, Julio

func main () {

// Opción 1
	
	var months = map[int]string{1:"Enero", 2:"Febrero", 3:"Marzo", 4:"Abril", 5:"Mayo", 6:"Junio", 7:"Julio", 8:"Agosto", 9:"Septiembre", 10:"Octubre", 11:"Noviembre", 12:"Diciembre"}

	number := 8
	fmt.Printf(months[number])


// Opcion 2

//    month := 7
//    switch month {
// 		case 1:
// 			fmt.Println("Enero")
// 		case 2:
// 			fmt.Println("Febrero")
// 		case 3:
// 			fmt.Println("Marzo")
// 		case 4:
// 			fmt.Println("Abril")
// 		case 5:
// 			fmt.Println("Mayo")
// 		case 6:
// 			fmt.Println("Junio")
// 		case 7:
// 			fmt.Println("Julio")
// 		case 8:
// 			fmt.Println("Agosto")
// 		case 9:
// 			fmt.Println("Septiembre")
// 		case 10:
// 			fmt.Println("Octubre")
// 		case 11:
// 			fmt.Println("Noviembre")
// 		case 12:
// 			fmt.Println("Diciembre")	
// 		default:
// 			fmt.Println("Desconocido")
//    }
	

}