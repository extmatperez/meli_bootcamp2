package main

import "fmt"

func main() {

	/* Ejercicio 1 */
	nombre_1 := "Matias"
	apellido_1 := "De Bonis"

	fmt.Println(nombre_1,apellido_1)

	/* Ejercicio 2 */
	var temperatura float64
	temperatura = 19.0
	var humedad float64
	humedad = 35.0
	var presion int
	presion = 1024

	fmt.Printf("Temperatura: %v, Humedad %v%%, presión: %v", temperatura,humedad, presion)

	/* Ejercicio 3 */
	var nombre string /* Las variables no pueden comenzar con un numero */
	var apellido string
	var edad int /* Primero se declara el nombre de la variable y luego el tipo */
	apellido = "Apellido" /* Las variables no pueden comenzar con un numero, la variable ya esta declarda y los : no son necesarios, la variable esta declarada como string */
	var licencia_de_conducir bool = true
	var estatura_de_la_persona int /* Las variables no pueden contener espacios */
	cantidadDeHijos := 2

	fmt.Println(nombre, apellido, edad, licencia_de_conducir, estatura_de_la_persona, cantidadDeHijos)

	/* EJercicio 4 */

	var apellido_4 string = "Gomez"
	var edad_4 int = 35
	boolean := false
	var sueldo float64 = 45857.90
	var nombre_4 string = "Julián"

	fmt.Println(apellido_4, edad_4, boolean, sueldo, nombre_4)
  
  
}