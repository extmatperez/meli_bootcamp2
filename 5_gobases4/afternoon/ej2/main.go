package main

import (
	"fmt"
	"os"
)

type badDataError struct {
	fields string
}

func (bde *badDataError) Error() string {
	return fmt.Sprintf("the fields %s are incorrect", bde.fields)
}

type client struct {
	file     int
	fullName string
	document int
	phone    int
	address  string
}

func generateID() *int {
	return nil
}

func createClient() client {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	file := generateID()
	if file == nil {
		panic("id not exists")
	}

	return client{
		*file,
		"Archuby Federico",
		654654,
		65634643,
		"La Plata",
	}
}

func openFile(file string) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err := os.ReadFile(file)
	if err != nil {
		panic("file not found or broken")
	}
}

func checkInt(value int, msg string, valid bool, name string) (bool, string) {
	if value == 0 {
		return false, fmt.Sprintf("%s %s", msg, name)
	}

	return valid, ""
}

func checkString(value string, msg string, valid bool, name string) (bool, string) {
	if value == "" {
		return false, fmt.Sprintf("%s %s", msg, name)
	}

	return valid, ""
}

func checkClient(newClient client) (bool, error) {
	msg := ""
	valid := true

	valid, msg = checkInt(newClient.file, msg, valid, "file")
	valid, msg = checkInt(newClient.document, msg, valid, "document")
	valid, msg = checkInt(newClient.phone, msg, valid, "phone")

	valid, msg = checkString(newClient.fullName, msg, valid, "fullName")
	valid, msg = checkString(newClient.address, msg, valid, "address")

	if msg != "" {
		return false, &badDataError{
			fields: msg,
		}
	}

	return true, nil
}

func main() {
	defer func() {
		fmt.Println("No files remained open")
	}()

	defer func() {
		fmt.Println("Various errors were detected in execution time")
	}()

	newClient := createClient()

	openFile("../customers.txt")

	_, err := checkClient(newClient)

	if err != nil {
		fmt.Println(err)
	}

	println("Fin de la ejecución")
}
