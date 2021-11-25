package main

import (
	"fmt"
	"os"
)

func main() {

	var err error
	err = os.Setenv("NAME", "matias")

	variable := os.Getenv("NAME")
	value2, ok := os.LookupEnv("NAME")

	if ok {
		fmt.Printf("\nSe encontro la variable de entorno: %s", value2)
	} else {
		fmt.Printf("\nNo encontro la variable de entorno: %s", value2)
	}
	fmt.Println(variable)
	fmt.Println(err)

	data, err := os.ReadFile("./archivo/myFile.txt")
	if err == nil {
		file := string(data)
		fmt.Println(file)

	} else {
		fmt.Println("El archivo no existe...")
	}

}
