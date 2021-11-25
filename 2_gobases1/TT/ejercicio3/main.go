package main

import "fmt"

type Persona struct {
	Edad       int
	EsEmpleado bool
	Antiguedad int //solo es valido si EsEmpleado == true
	Sueldo     int //solo es valido si EsEmpleado == true
}

func main() {
	//persona := cargarPersona()
	//calcularPrestamo(persona)

	fmt.Println("Persona1: ")
	p1 := Persona{15, false, 0, 0}
	calcularPrestamo(p1)
	fmt.Println("Persona2: ")
	p2 := Persona{15, true, 4, 0}
	calcularPrestamo(p2)
	fmt.Println("Persona3: ")
	p3 := Persona{25, false, 0, 0}
	calcularPrestamo(p3)
	fmt.Println("Persona4: ")
	p4 := Persona{25, true, 0, 0}
	calcularPrestamo(p4)
	fmt.Println("Persona5: ")
	p5 := Persona{25, true, 5, 0}
	calcularPrestamo(p5)
	fmt.Println("Persona6: ")
	p6 := Persona{25, true, 5, 200000}
	calcularPrestamo(p6)

}

func cargarPersona() Persona {
	var edad, antiguedad, sueldo int
	var empleada string
	fmt.Printf("Ingrese la edad de la persona: ")
	fmt.Scanf("%d", &edad)
	for empleada != "Y" && empleada != "N" {
		fmt.Printf("¿La persona está actualmente empleada?: (Y/N)")
		fmt.Scanf("%s", &empleada)
		fmt.Println(empleada)
	}
	if empleada == "Y" {
		fmt.Printf("Ingrese la antigüedad en años de la persona en el trabajo: ")
		fmt.Scanf("%d", &antiguedad)
		fmt.Printf("Ingrese el sueldo de la persona: $")
		fmt.Scanf("%d", &sueldo)
	} else {
		antiguedad, sueldo = 0, 0
	}
	return Persona{edad, empleada == "Y", antiguedad, sueldo}
}

func calcularPrestamo(persona Persona) {
	/*Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años
	se encuentren empleados
	tengan más de un año de antigüedad en su trabajo.

	Dentro de los préstamos que otorga no les cobrará interés
		a los que su sueldo es mejor a $100.000.
	*/

	switch {
	case persona.Edad <= 22:
		fmt.Printf("Solo es posible otorgar préstamos a mayores de 22 años.")
	case !persona.EsEmpleado || persona.Antiguedad < 1:
		fmt.Printf("Solo es posible otorgar préstamos a personas empleadas hace más de un año.")
	case persona.Sueldo > 100000:
		fmt.Printf("Puedes acceder al préstamo sin interés")
	default:
		fmt.Printf("Puedes acceder al préstamo con interés")
	}
	fmt.Println()
}
