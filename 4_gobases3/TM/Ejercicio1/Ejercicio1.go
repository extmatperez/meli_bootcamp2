package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const file = "/Users/rovega/Documents/GitHub/meli_bootcamp2/4_gobases3/TM/Ejercicio1/a.txt"

func Write(p []Producto) {
	var textoFormateado string

	for _, element := range p {
		prodFormateado, err := json.Marshal(element)
		// Esta diferenciacion es porque el ultimo no deberia tener ; porque sino falla la lectura.
		if err == nil && element != p[len(p)-1] {
			textoFormateado += fmt.Sprintf("%v;", string(prodFormateado))
		}
		if err == nil && element == p[len(p)-1] {
			textoFormateado += fmt.Sprintf("%v", string(prodFormateado))
		}
	}

	archivo, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Algo salio mal al abrir el archivo.")
	}

	defer archivo.Close()

	if _, err = archivo.WriteString(textoFormateado); err != nil {
		fmt.Println("Algo salio mal al escribir el archivo.")
	} else {
		fmt.Println("Se guardo correctamente.")
	}
}

func Read() {
	var products []string
	data, err3 := os.ReadFile(file)
	if err3 == nil {
		products = strings.Split(string(data), ";")
		for _, element := range products {
			var pr = &Producto{}
			err := json.Unmarshal([]byte(element), pr)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%d\t\t%.2f\t%d\n", pr.Id, pr.Precio, pr.Cantidad)
		}
	} else {
		fmt.Println("No se encuentra archivo.")
	}
}

func Total_Calculate() float64 {
	var products []string
	var total float64
	data, err3 := os.ReadFile(file)
	if err3 == nil {
		products = strings.Split(string(data), ";")
		for _, element := range products {
			var pr = &Producto{}
			err := json.Unmarshal([]byte(element), pr)
			if err != nil {
				panic(err)
			}
			total = total + (pr.Precio * (float64)(pr.Cantidad))
		}
	} else {
		fmt.Println("No se encuentra archivo.")
		return 0
	}
	return total
}

type Producto struct {
	Id       int     `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

func main() {
	p0 := Producto{
		Id:       1,
		Precio:   245.67,
		Cantidad: 7}
	p1 := Producto{
		Id:       2,
		Precio:   45.67,
		Cantidad: 15}
	p2 := Producto{
		Id:       3,
		Precio:   185.67,
		Cantidad: 2}
	p3 := Producto{
		Id:       4,
		Precio:   465.67,
		Cantidad: 4}

	Write([]Producto{p0, p1, p2, p3})
	fmt.Printf("%s\t\t%s\t%s\n", "ID", "Precio", "Cantidad")
	Read()
	fmt.Printf("\t\t%.2f\t", Total_Calculate())
}
