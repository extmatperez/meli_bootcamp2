// Tambien es el ejercicio 2. Poes paja de hacer otro file.
// Completar porque no esta funcionando.
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const file = "/Users/rovega/Documents/GitHub/meli_bootcamp2/5_gobases4/TT/customers.txt"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Customer struct {
	Legajo    string `json:"legajo"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	DNI       string `json:"dni"`
	Telefono  string `json:"telefono"`
	Direccion string `json:"direccion"`
}

func Read() {
	var customers []string
	data, err3 := os.ReadFile(file)

	if err3 != nil {
		defer func() {
			fmt.Println("Ejecución finalizada!")
		}()
	}

	if err3 == nil {
		customers = strings.Split(string(data), ";")
		for _, element := range customers {
			var pr = Customer{}
			err := json.Unmarshal([]byte(element), pr)

			defer func() {
				fmt.Println("Ejecución finalizada!")
			}()

			if err != nil {
				defer func() {
					fmt.Println(pr)
				}()
				panic("El archivo no fue encontrado o se encuentra dañado. Si")
			}
			fmt.Printf("%s\t\t%s %s\t%s\t%s\t%s", pr.Legajo, pr.Nombre, pr.Apellido, pr.DNI, pr.Telefono, pr.Direccion)
		}
	} else {
		panic("El archivo no fue encontrado o se encuentra dañado.")
	}
}

func generar_legajo(n int) (string, error) {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	if b == nil {
		panic("No se pudo generar legajo de cliente.")
	}
	return string(b), nil
}

func verificar_existencia(dni string) {
	var existing_customers []string
	data, err3 := os.ReadFile(file)

	if err3 == nil {
		existing_customers = strings.Split(string(data), ";")
		for _, element := range existing_customers {
			var pr = &Customer{}
			err0 := pr.DNI == dni
			err := json.Unmarshal([]byte(element), pr)

			if err0 != false {
				panic("Existe un usuario con ese DNI.")
			}

			if err != nil {
				panic(err)
			}

			fmt.Printf("%s\t\t%s %s\t%s\t%s\t%s", pr.Legajo, pr.Nombre, pr.Apellido, pr.DNI, pr.Telefono, pr.Direccion)
		}
	} else {
		panic("El archivo no fue encontrado o se encuentra dañado.")
	}
}

func main() {
	Read()
	var legajo, dni /*nombre, apellido, telefono, direccion*/ string
	legajo, err0 := generar_legajo(8)
	if err0 != nil {
		panic("No se pudo generar legajo de cliente.")
	}
	fmt.Println("Ingrese el DNI:")
	fmt.Scanf("%s", &dni)
	print(legajo, dni)
}
