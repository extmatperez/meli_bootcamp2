/* var 1nombre string  ---------- var nombre1 string (no debe iniciar con número)
  var apellido string
  var int edad   ------------ var edad int
  1apellido := 6 ------------ apellido1 := (no debe iniciar con número y debería ser un string)
  var licencia_de_conducir = true ---------- var licencia_de_conducir bool = true
  var estatura de la persona int ------------- var estatura_de_la_persona int
  cantidadDeHijos := 2 */

  /* var apellido string = "Gomez"
  var edad int = "35" --------- var edad int = 35 (sin "")
  boolean := "false"; --------- Es correcto.... también podría ser boolean := false (sin "")
  var sueldo string = 45857.90  var sueldo float64 = 45857.90
  var nombre string = "Julián" */

  package main

  import "fmt"

  func main(){
	var nombre1 string
	var apellido string
	var edad int
	apellido1 := 6
	var licencia_de_conducir = true
	var estatura_de_la_persona int
	cantidadDeHijos := 2
	/* ---------------------- */
	var edad2 int = 35
	boolean := "false"
	var sueldo float64 = 45857.90
	var nombre string = "Julián"

	fmt.Println("Nombre: ", nombre1)
	fmt.Println("Apellido: ", apellido)
	fmt.Println("Edad: ", edad)
	fmt.Println("Apellido1: ", apellido1)
	fmt.Println("Licencia de conducir: ", licencia_de_conducir)
	fmt.Println("Estatura de la persona : ", estatura_de_la_persona)
	fmt.Println("Cantidad de hijos: ", cantidadDeHijos)
	fmt.Println("Edad: ", edad2)
	fmt.Println("Boolean: ", boolean)
	fmt.Println("Sueldo: ", sueldo)
	fmt.Println("Nombre: ", nombre)
  }
