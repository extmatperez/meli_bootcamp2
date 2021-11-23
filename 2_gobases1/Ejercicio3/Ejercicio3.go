package main

import "fmt"

func main() {
	var lnombre string = "Pablo"

	var apellido string = "Perez"

	//var int edad -- Aca el orden esta invertido y debe darse vuelta
	var edad int = 2

	lapellido := 6 // Es correcta la asignación, pero no corresponde con el nombre antes definido (apellido) ni su tipo string.

	//var licencia_de_conducir = true -- Falta boolean, sino no sabe que tipo es.
	licencia_de_conducir := true // o var licencia_de_conducir boolean = true

	//var estatura de la persona int -- Aca se toman varios nombres, cuando deberia ser snake o camel case. Ademas de que la estatura se debe tratar de un flotante.
	var estatura_persona float64 = 1.82

	//cantidadDeHijos := 2 -- Si todo el resto de las variables se determino que sea con snake case, esta no puede ser camel case, no queda bien.
	cantidad_hijos := 2

	fmt.Println(lnombre, apellido, edad, lapellido, licencia_de_conducir, estatura_persona, cantidad_hijos)
}
