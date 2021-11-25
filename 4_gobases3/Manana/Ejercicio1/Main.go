package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Compras struct {
	Registro []Compra `json:"total_compras"`
}

type Compra struct {
	Id       string `json:"id_producto"`
	Precio   string `json:"precio_producto"`
	Cantidad string `json:"cantidad_producto"`
}

func (c *Compras) nuevaCompra(compra Compra) {

	/*
		res, err := os.ReadFile("productos.txt")
		if err != nil {
			fmt.Println(err)
		}
	*/

	fmt.Println("ANTES DE AGREGAR", c.Registro)
	c.Registro = append(c.Registro, compra)
	fmt.Println("DESPUES DE AGREGAR", c.Registro)

	comprasJSON, errCJson := json.Marshal(c.Registro)
	if errCJson != nil {
		fmt.Println(errCJson)
	} else {
		stringCJSON := string(comprasJSON)
		errUpdate := os.WriteFile("productos.txt", []byte(stringCJSON), 0644)
		if errUpdate != nil {
			fmt.Println(errUpdate)
		} else {
			fmt.Println("Success")
		}
	}

}

func main() {

	fmt.Println("Bienvenidos al ejercicio 1")

	/*
		archivos, err := os.ReadDir(".")

		if err != nil {
			fmt.Println(err)
		} else {
			for _, archivo := range archivos {
				fmt.Println(archivo)
			}
		}
	*/

	/*
		texto, err := os.ReadFile("productos.txt")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(texto))
		}
	*/

	/*
		producto1 := Compra{"16", "1500.0", "2"}
		producto2 := Compra{"54", "1800", "9"}
		fmt.Println("PRODUCTO1", producto1)
	*/

	/*
		producto1JSON, errJSON := json.Marshal(producto1)
		if errJSON != nil {
			fmt.Println(errJSON)
		}
		fmt.Println("JSON", string(producto1JSON))

		producto1JSON1 := string(producto1JSON)
		fmt.Println("Nuevo JSON", producto1JSON1)
		errWrite := os.WriteFile("productos.txt", []byte(producto1JSON1), 0644)
		if errWrite != nil {
			fmt.Println(errWrite)
		}

		texto, errRead := os.ReadFile("productos.txt")
		if errRead != nil {
			fmt.Println(errRead)
		} else {
			fmt.Println("El read", string(texto))
		}
	*/

	/*
		Compras1 := Compras{}
		Compras1.nuevaCompra(producto1)
		Compras1.nuevaCompra(producto2)
	*/

	texto, err := os.ReadFile("productos.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Paso0", string(texto))
		//bufio.NewScanner(b)
		paso1 := strings.NewReader(string(texto))
		fmt.Println("Paso1", paso1)
		paso2, errPaso2 := io.ReadAll(paso1)
		if errPaso2 != nil {
			fmt.Println(errPaso2)
		} else {
			fmt.Println("Paso2", paso2)
			paso3 := strings.NewReader(string(paso2))
			fmt.Println("Paso3", paso3)
			paso4 := bufio.NewScanner(paso3)
			fmt.Println("Paso4", paso4)
		}
	}

}
