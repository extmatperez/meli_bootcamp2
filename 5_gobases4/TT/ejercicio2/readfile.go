package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func LeerArchivo() (customers []Customer, err error) {
	defer func() {
		er := recover()
		if er != nil {
			fmt.Println("error:", er)
		}

	}()

	data, err := os.ReadFile("./" + os.Getenv("ARCHIVO"))

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		err = json.Unmarshal(data, &customers)
	}

	return
	//mt.Printf("File read.\n")
}
