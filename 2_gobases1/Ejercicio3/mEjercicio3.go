package main
import "fmt"


// Los comentarios encima de las declaraciones indican que dichas sentacias son corregidas
// en la linea siguiente


func main() {
//var 1nombre string
var nombre string;
var apellido string;
//var int edad
var edad int
//apellido := 6
apellido = "PEREZ"

// var licencia_de_cconducir = true
var licenciaDeCconducir = true

//var estatura de la persona int
var estaturaDeLaPersona float32 = 1.80

cantidadDeHijos := 2
fmt.Println("El nombre completo de la persona es:", nombre," " ,apellido) ;
fmt.Println("Edad: ",edad );
fmt.Println("¿Puede conducir?", licenciaDeCconducir)
fmt.Println("Estatura: ", estaturaDeLaPersona)
fmt.Println("¿Cuantos hijos tiene?", cantidadDeHijos)
  
}
