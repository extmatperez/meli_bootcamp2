package main
import "fmt"





func main() {
// Los comentarios encima de las declaraciones indican que dichas sentacias son corregidas
// en la linea siguiente

var apellido string = "Gomez"

//	var edad int = "35" 
	var edad int = 35

//	boolean := "false";
	boolean  := false;

//	var sueldo string = 45857.90
	var sueldo float32 = 45857.90

var nombre string = "Julián"
  
fmt.Println("El nombre completo de la persona es:", nombre," " ,apellido) ;
fmt.Println("Edad: ",edad );
fmt.Println("¿Puede conducir?", boolean)
fmt.Println("Sueldo: ", sueldo)

  
}
