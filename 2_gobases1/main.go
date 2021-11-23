package main

import "fmt"

func main() {

	//Ejercicio 1
	var nombrePropio, dir string = "Federico Archuby", "La Plata"

	fmt.Printf("Mi nombre es %v y vivo en %v\n", nombrePropio, dir)

	//Ejercicio 2

	/*
	 * Cómo la temperatura suele utilizarse un decimal, decidí utilizar un flotante. La presión se suele
	 *
	 * La presión se suele representar en hectopascales. Al medirse con la medida de hecto, se suele utilizar la coma.
	 *     Utilice como dato el valor de la presión sobre el nivel del mar.
	 *
	 * La humedad se suele representar en porcentaje y sin decimales.
	 */
	var temp, presion float32 = 21.5, 1013.25

	var humedad = 47 //
	fmt.Printf("Hoy en %v hacen %.1f grados, con una humedad de %d%% y una presión de %.2f hPa\n", dir, temp, humedad, presion)

	//Ejercicoi 3
	var nombre string   //Incorrecta, el nombre de la variable empezaba con un número.
	var apellido string //Correcta
	var edad int        //Incorrecta, el nombre debe ir antes del tipo.
	promedio := 6
	/* Incorrecta, el nombre de la variable empezaba con un número.
	 * A la hora de editar se debe tener en cuenta que el nombre no coincida con la variable apellido ya declarada. Además el nombre no es descriptivo.
	 */
	var licencia_de_conducir bool = true //Incorrecta, falta el tipo de dato
	var estaturaDeLaPersona int          //Incorrecta, no pueden haber espacios en el nombre de la variable.
	cantidadDeHijos := 2                 //Correcta

	fmt.Println("Imprimo las variables para usarlas", nombre, apellido, edad, promedio, licencia_de_conducir, estaturaDeLaPersona, cantidadDeHijos)

	//Ejercicio 4
	//Cambio el nombre de algunas variables para que no se repitan de las declaradas en otros ejercicios.
	var apellidoEj3 string = "Gomez" //Correcta
	var edadEj3 int = 35             //Incorrecta, el dato a cargar no era un entero sino un string.
	recibido := "false"              //Incorrecta, al usar el := se debe decir directamente el nombre de la variable y no el tipo de dato. Este último se difiere. Además el ; no se usa.
	//Si bien el nombre boolean no es la palabra reservada para el tipo de dato(bool),
	//no es correcto que la variable de información sobre su tipo en el nombre, sino que debe decir para que se va a usar.
	var sueldo float32 = 45857.90   //Incorrecta, se intentaba guardar un flotante en un string. El tipo de dato a usar estaba mal.
	var nombreEj3 string = "Julián" // Correcta.

	fmt.Println("Imprimo las variables para usarlas", apellidoEj3, edadEj3, recibido, sueldo, nombreEj3)

}
