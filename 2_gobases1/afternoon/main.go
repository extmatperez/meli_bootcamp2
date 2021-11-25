package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	//Ejercicio 1
	reader := bufio.NewReader(os.Stdin)
	var palabra string
	fmt.Println("Ingrese la palabra: ")
	palabra, _ = reader.ReadString('\n')
	var letras []string = strings.Split(palabra, "")
	fmt.Printf("La palabra tiene %d letras\n", len(letras))
	fmt.Print("Palabra deletreada: ")
	for _, letra := range letras {
		fmt.Printf("%s ", letra)
	}
	fmt.Println()

	//Ejercicio 2
	fmt.Println("")
	var precio, descuento float32
	fmt.Println("Ingrese el precio: ")
	fmt.Scanf("%f", &precio)
	fmt.Println("Ingrese el descuento en porcentaje: ")
	fmt.Scanf("%f", &descuento)

	precioFinal := precio - ((precio * descuento) / 100)
	fmt.Printf("El precio final es de %v pesos\n", precioFinal)

	//Ejercicio 3
	fmt.Println("")
	var (
		edad, antiguedad int
		sueldo           float32
		empleado         bool
	)

	fmt.Println("Ingrese la edad: ")
	fmt.Scanf("%d", &edad)
	fmt.Println("Ingrese la antigüedad: ")
	fmt.Scanf("%d", &antiguedad)
	fmt.Println("Ingrese el sueldo: ")
	fmt.Scanf("%f", &sueldo)
	rand.Seed(time.Now().UnixNano())
	empleado = rand.Float32() < 0.5
	fmt.Printf("Random de empleado salio %v\n", empleado)
	if (empleado) && (edad > 22) && (antiguedad > 1) {
		if sueldo > 100000 {
			fmt.Println("Se le otorgará el crédito sin intereses")
		} else {
			fmt.Println("Se le otorgará el crédito con intereses")
		}
	} else {
		fmt.Println("No se le otorgará el crédito")
	}

	//Ejercicio 4
	fmt.Println("")
	var meses = [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	var mes int
	fmt.Println("Ingrese el número del mes: ")
	fmt.Scanf("%d", &mes)
	fmt.Printf("El mes ingresado fue %s\n", meses[mes-1])

	//Ejercicio 5
	fmt.Println("")
	var estudiantes = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Printf("Los estudiantes son %v\n", estudiantes)
	estudiantes = append(estudiantes, "Gabriela")
	fmt.Printf("Los estudiantes son %v\n", estudiantes)

	//Ejercicio 6
	fmt.Println("")
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("Benjamín tiene %d años\n", employees["Benjamin"])

	var cantidad int = 0
	for _, edad := range employees {
		if edad >= 21 {
			cantidad++
		}
	}
	fmt.Printf("Hay %d empleados con mas de 21 años\n", cantidad)
	fmt.Printf("Los empleados son %v\n", employees)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Printf("Los empleados son %v\n", employees)
}
