package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"os"
	"strconv"
)

const (
	fileName = "customers.txt"
)

type cliente struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Legajo    string `json:"legajo"`
	DNI       string `json:"dni"`
	Telefono  string `json:"telefono"`
	Domicilio string `json:"domicilio"`
}

func generateID(dni string) string {
	if dni == "" || dni == "\n" {
		return ""
	}
	algorithm := fnv.New32a()
	algorithm.Write([]byte(dni))
	s := strconv.FormatUint(uint64(algorithm.Sum32()), 10)
	return s
}

func readFile(fileName string) []byte {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("error: ", err)
			c := make(map[string]cliente)
			b, _ := json.Marshal(c)
			os.WriteFile(fileName, b, 0644)
		}
	}()

	content, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		panic("el archivo indicado no fue encontrado o está dañado\n")
	}
	return content
}

func loadInformation(c []byte) map[string]cliente {

	clientsMap := make(map[string]cliente)

	json.Unmarshal(c, &clientsMap)

	return clientsMap
}

func verifyData(name, lastName, address, phone, dni string) (cliente, error) {

	if name == "" {
		return cliente{}, errors.New("El nombre es obligatorio")
	}
	if lastName == "" {
		return cliente{}, errors.New("El apellido es obligatorio")
	}
	if address == "" {
		return cliente{}, errors.New("La dirección es obligatoria")
	}
	if phone == "" {
		return cliente{}, errors.New("El telefono es obligatorio")
	}

	c := cliente{Nombre: name, Apellido: lastName, DNI: dni, Telefono: phone, Domicilio: address}

	return c, nil

}

func saveOnFile(c map[string]cliente) {

	fileString, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error marshaliando el JSON")
	}

	os.WriteFile(fileName, fileString, 0644)
}

func addClient(c map[string]cliente) (bool, error) {
	var name string
	var lastName string
	var dni string
	var address string
	var phone string
	var id string
	fmt.Println("Agregar cliente: ")

	fmt.Printf("DNI: ")
	fmt.Scanf("%s", &dni)
	id = generateID(dni)
	if id == "" {
		panic("Error generando id de legajo")
	}
	if c[id].Legajo != "" {
		return false, fmt.Errorf("Error el usuario con DNI %s ya exitse", dni)
	}

	fmt.Printf("Nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Apellido: ")
	fmt.Scanf("%s", &lastName)

	fmt.Printf("Dirección: ")
	fmt.Scanf("%s", &address)
	fmt.Printf("Telefono: ")
	fmt.Scanf("%s", &phone)

	client, err := verifyData(name, lastName, address, phone, dni)
	if err != nil {
		return false, errors.New("Error existen datos nulos")
	}

	client.Legajo = id
	c[id] = client
	saveOnFile(c)

	return false, nil

}

func main() {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Println("Fin de la ejecución")
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		fmt.Println("No han quedado archivos abiertos")

	}()

	clientsMap := loadInformation(readFile(fileName))

	for {
		var option int
		fmt.Println("Menu: \n1. Mostrar lista. \n2. Agregar cliente.")
		fmt.Scanf("%d", &option)
		if option == 1 {
			fmt.Printf("\n%+v\n", clientsMap)
		} else if option == 2 {
			isAdd, err := addClient(clientsMap)
			if err != nil {
				fmt.Println("Error agregando el cliente: ", err)
			}
			if isAdd == true {
				fmt.Println("Cliente agregado correctamente")
			}
		} else {
			fmt.Println("Opcion invalida")
		}
	}

}
