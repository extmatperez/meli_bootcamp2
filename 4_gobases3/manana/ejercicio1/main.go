package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type producto struct {
	Id       string `json:"id"`
	Precio   string `json:"precio"`
	Cantidad string `json:"cantidad"`
}

func guardarArch() {

	pr1 := &producto{Id: "147", Precio: "123", Cantidad: "3"}
	pr1Doc, _ := json.Marshal(pr1)

	f, err := os.OpenFile("./myfile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("error :", err)
	}

	defer f.Close()

	if _, err = f.WriteString(string(pr1Doc)); err != nil {
		fmt.Println("error :", err)
	}
}

func main() {
	guardarArch()
}
