package main

import (
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

	f, err := os.OpenFile("./myfile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("error :", err)
	}

	defer f.Close()

	if _, err = f.WriteString(pr1.Id + ";" + pr1.Precio + ";" + pr1.Cantidad + "\n"); err != nil {
		fmt.Println("error :", err)
	}
}

func main() {
	guardarArch()
}
