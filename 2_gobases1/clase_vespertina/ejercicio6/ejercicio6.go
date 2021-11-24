package main

import (
	"fmt"
)

var empleados = map[string]uint8{
	"Benjamin": 20,
	"Nahuel":   26,
	"Brenda":   19,
	"Dario":    44,
	"Pedro":    30,
}

func main() {
	imprimirMenu()
}

func imprimirMenu() {
	fmt.Println("Elija una opcion:")
	fmt.Println("1. Imprimir edad empleado")
	fmt.Println("2. Obtener empleados mayores a 21")
	fmt.Println("3. Agregar empleado")
	fmt.Println("4. Eliminar empleado")
	fmt.Println("5. Salir")
	var opcion uint8
	fmt.Println("Ingrese una opcion:")
	fmt.Scanf("%d", &opcion)
	fmt.Println("\n")
	procesarOpcion(opcion)
}

func procesarOpcion(opcion uint8) {
	if opcion < 1 || opcion > 5 {
		fmt.Println("Opcion incorrecta")
		fmt.Println("\n")
		imprimirMenu()
	}

	switch {
	case opcion == 1:
		imprimirEdad()
	case opcion == 2:
		obtenerMayores21()
	case opcion == 3:
		agregarEmpleado()
	case opcion == 4:
		eliminarEmpleado()
	case opcion == 5:
		break
	}
}

func imprimirEdad() {
	fmt.Println("Ingrese el nombre del empleado a consultar:")
	var nombreEmpleado string
	fmt.Scanf("%v", &nombreEmpleado)
	fmt.Println("\n")
	if _, ok := empleados[nombreEmpleado]; ok {
		fmt.Printf("Empleado: %v - Edad: %v", nombreEmpleado, empleados[nombreEmpleado])
		fmt.Println("\n")
	} else {
		fmt.Println("El empleado no existe.")
		fmt.Println("\n")
	}
	imprimirMenu()
}

func obtenerMayores21() {
	for clave, valor := range empleados {
		if empleados[clave] > 21 {
			fmt.Printf("Empleado: %v - Edad: %d", clave, valor)
			fmt.Println("")
		}
	}
	fmt.Println("")
	imprimirMenu()
}

func agregarEmpleado() {
	fmt.Printf("Ingrese el nombre del empleado:")
	var nombreEmpleado string
	fmt.Scanf("%v", &nombreEmpleado)
	fmt.Println("")
	fmt.Printf("Ingrese la edad del empleado:")
	var edadEmpleado uint8
	fmt.Scanf("%d", &edadEmpleado)
	fmt.Println("")
	empleados[nombreEmpleado] = edadEmpleado
	fmt.Printf("Se agrego el empleado %v", nombreEmpleado)
	fmt.Printf("\n")
	imprimirMenu()
}

func eliminarEmpleado() {
	fmt.Printf("Ingrese el nombre del empleado:")
	var nombreEmpleado string
	fmt.Scanf("%v", &nombreEmpleado)
	fmt.Println("")
	delete(empleados, nombreEmpleado)
	fmt.Printf("Se ha eliminado el empleado %v", nombreEmpleado)
	fmt.Println("\n")
	imprimirMenu()
}
