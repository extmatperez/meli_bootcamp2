// Tambien es el ejercicio 2. Poes paja de hacer otro file.
// Completar porque no esta funcionando.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Client struct {
	FileNumber  int64  `json:"file_number"`
	Fullname    string `json:"fullname"`
	DNI         int    `json:"dni"`
	PhoneNumber int    `json:"phone_number"`
	Address     string `json:"address"`
}

func generateFileNumber() int64 {
	return time.Now().Unix()
}

func clientExists(newFileNumber int64, clients []Client) bool {
	for i := 0; i < len(clients); i++ {
		if clients[i].FileNumber == newFileNumber {
			return true
		}
	}
	return false
}

func newClient(newFileNumber int64, fullname string, dni int, phoneNumber int, address string) (Client, error) {
	return Client{
		FileNumber:  newFileNumber,
		Fullname:    fullname,
		DNI:         dni,
		PhoneNumber: phoneNumber,
		Address:     address,
	}, nil
}

func validateFieldsOfClient(fileNumber int64, fullname string, dni, phoneNumber int, address string) (bool, error) {
	if fileNumber != 0 && fullname != "" && dni != 0 && phoneNumber != 0 && address != "" {
		return true, nil
	}

	if fileNumber == 0 {
		return false, &ClientError{
			Status:  500,
			Message: "El número de legajo no puede ser 0",
		}
	}

	if fullname == "" {
		return false, &ClientError{
			Status:  500,
			Message: "El nombre completo no puede estar vacio",
		}
	}

	if dni == 0 {
		return false, &ClientError{
			Status:  500,
			Message: "El DNI no puede ser 0",
		}
	}

	if phoneNumber == 0 {
		return false, &ClientError{
			Status:  500,
			Message: "El número de teléfono no puede ser 0",
		}
	}

	if address == "" {
		return false, &ClientError{
			Status:  500,
			Message: "El domicilio no puede estar vacio",
		}
	}

	return false, &ClientError{
		Status:  500,
		Message: "Error desconocido",
	}
}

func readCustomersFile() []Client {
	bytes, err := os.ReadFile("./customers.txt")

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)

		}
	}()

	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	}

	clients := []Client{}
	json.Unmarshal(bytes, &clients)

	return clients
}

func writeCustomersFile(clients []Client) error {
	clientsJson, _ := json.Marshal(clients)

	errWriteFile := os.WriteFile("./customers.txt", clientsJson, 0644)

	if errWriteFile != nil {
		return errors.New("error escribiendo el archivo con clientes")
	}

	return nil
}

type ClientError struct {
	Status  int
	Message string
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("status: %d - message: %s", e.Status, e.Message)
}

func main() {
	var clients []Client

	clients = readCustomersFile()

	newFileNumber := generateFileNumber()

	if newFileNumber < 0 { // int cannot be nil, thats why i check if newFileNumber is 0 or negative
		panic("El número de legajo no puede ser nil")
	}

	clientExists := clientExists(newFileNumber, clients)

	if !clientExists {
		fullname := "Nicolas Ziliotto"
		dni := 38296195
		phoneNumber := 2302489894
		address := "Calle 30 Nº 964"

		validated, err := validateFieldsOfClient(newFileNumber, fullname, dni, phoneNumber, address)

		if err != nil || !validated {
			fmt.Println("No se pudo validar los datos del cliente")
			fmt.Println(err)
		}

		if validated {
			client, _ := newClient(newFileNumber, fullname, dni, phoneNumber, address)
			clients = append(clients, client)

			err := writeCustomersFile(clients)

			if err != nil {
				fmt.Println(err)
			}
		}
	}

	fmt.Println("Fin de la ejecucion")
}
