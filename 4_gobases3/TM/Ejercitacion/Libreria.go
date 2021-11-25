package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// n para salto de linea
// \\ barra invertida
// t para tabulacion
// Verbos de impresion: Todos acompañados por % - v (valor en formato estandar)
// T tipo de dato del valor a imprimir, t es un bool
// s un string, f un float, d un decimal, b un entero binario (muestra 1 byte), o un octal (muestra como 000 en escala octal, es como 8 al cuadrado, 8 a la 1 y 8 a la 0),
// c para caracteres, p una direccion de memoria.
// Siempre que precisemos una posicion en memoria tenemos que colocar & antes de la variable.

func main() {
	// Tenemos la funcion Sprint que concatena strings formando una nueva variable y luego mostrarla directo a la misma.
	var anio int = 25
	var nombre string = "Rodrigo"
	res := fmt.Sprintf("Mi nombre es %s y tengo %d años.", nombre, anio)
	fmt.Println(res)

	// PAQUETE OS.
	// Setenv. Funcion para definir variables de entorno. Si existe, la sobreescribe. Despues vemos como validarlo.
	err := os.Setenv("NAME", "gopher")
	// Getenv. Funcion para obtener variables de entorno.
	value := os.Getenv("NAME")
	fmt.Println(err, value)
	// LookupEnv. Equivalente a Getenv pero retorna dos valores: El valor en si y un booleano que indica si lo encontro o no (el "ok" que esta abajo)
	value_bis, ok := os.LookupEnv("NAME")
	if ok {
		fmt.Printf("Se encontro variable de entorno: %s\n", value_bis)
	} else {
		fmt.Println("No se encontro variable de entorno.")
	}
	// Exit. Hace que el programa termine de inmediato con el codigo de estado dado. El 0 indica exito, cualquier otro numero es un error. El resto no se ejecuta.
	//os.Exit(1)
	// ReadDir. Lee el directorio nombrado devolviendo las entradas.
	// ReadFile. Recibe como parametro el nombre del archivo absoluto y nos devuelve el contenido del archivo en bytes o un error si no lo encuentra.
	data, err := os.ReadFile("./myFile.txt")
	if err == nil {
		fmt.Println(string(data))
	} else {
		fmt.Println("No se encuentra archivo.")
	}

	// WriteFile. Escribe un archivo.
	err2 := os.WriteFile("./myFile.txt", []byte("Esto sume inicialmente."), 0644)

	if err2 != nil {
		fmt.Println("No funciono!")
	} else {
		fmt.Println("Funciono!")
	}

	data, err3 := os.ReadFile("./myFile.txt")
	if err3 == nil {
		fmt.Println(string(data))
	} else {
		fmt.Println("No se encuentra archivo.")
	}

	// PAQUETE IO
	// Copy. Toma un reader y hace un standard output.
	r := strings.NewReader("Algun io.Reader stream para ser leido\n")

	// ReadAll. Devuelve los datos que leyo y un error si lo hubiera.
	b, err := io.ReadAll(r)
	if err == nil {
		fmt.Println(b)
	} else {
		fmt.Println(err)
	}
	// WriteString. Escribe en un archivo.

}
