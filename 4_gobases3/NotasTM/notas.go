package main

// FMT
import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// FMT
	num := 15
	fmt.Print("")
	fmt.Printf("decimal: %d Ascii: %c octal: %o binario: %b hexa: %x %X valor memo: %v\n\n", num, num, num, num, num, num, &num)

	// OS

	err := os.Setenv("NAME", "gopher")
	//value := os.Getenv("NAME")
	value2, ok := os.LookupEnv("NAME")

	if ok {
		fmt.Printf("\n Se encontro la variable de entorno: %s", value2)
	} else {
		fmt.Printf("\n No se encontro la variable de entorno: %s", value2)
	}

	fmt.Println(err)
	//fmt.Println(value)
	fmt.Println("")

	fmt.Println("ReadFile")

	fmt.Println()
	data, err := os.ReadFile("./notas.go")
	if err != nil {
		fmt.Printf("\nEl archivo no se encontro.")
	} else {
		fmt.Printf("\nSe encontro el archivo: \n")
	}
	//fmt.Println(string(data))
	fmt.Println(err)

	//////////////////////////////////////////////

	fmt.Println()
	fmt.Printf("\nWriteFile()")

	p1 := persona{"Nacho", "Vargas"}
	p2 := persona{"Juan", "Perez"}

	// Coleccion de personas
	var lista []persona
	// appendeo persona
	lista = append(lista, p1, p2)
	// formateo json
	p1Formated, err := json.Marshal(lista)
	// Escribo el archivo con la nueva persona
	err = os.WriteFile("./archivo.txt", p1Formated, 0644)
	// Leo el archivo sobreescrito
	data, err = os.ReadFile("./archivo.txt")
	// creo una variable slice para unmarshalear el json
	var pListaLeida []persona
	// anmarshaleo el json
	json.Unmarshal(data, &pListaLeida)
	// traigo el nombre de la segunda persona
	fmt.Printf("\nEl nombre de la segunda persona es: %s", pListaLeida[1].Nombre)
	// printeo el json entero
	fmt.Printf("\n%+v", pListaLeida)

	if err == nil {
		file := string(data)
		fmt.Println()
		fmt.Println(file)
	} else {
		fmt.Printf("\nSe encontro el archivo: \n")
	}

	//fmt.Println(string(data))

	fmt.Println()
	fmt.Println()
	fmt.Println("Standard out put")

	fmt.Fprintln(os.Stdout, "\nG¡Hola")

}

type persona struct {
	Nombre   string `json: "nombre"`
	Apellido string `json: "apellido"`
}
