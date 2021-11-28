package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	Legajo    int
	Nombre    string
	Apellido  string
	DNI       int
	Telefono  int
	Domicilio string
}

type errorDetail struct {
	Nombre      string
	Descripcion string
}

func (e *errorDetail) Error() string {
	return fmt.Sprintf("%v:  %v", e.Nombre, e.Descripcion)
}

func generarLegajo() (int, error) {
	min := 100000
	max := 999999
	ID := rand.Intn(max-min) + min
	if ID <= min || ID >= max {
		return 1, nil
	}
	return ID, errors.New("Legajo creado correctamente")
}

func leerArchivo() {
	data, err := os.ReadFile("customers.txt")

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic("El archivo no fue creado o esta dañado")
	}

	fmt.Println(data)

}

func verifNum(l int) (int, error) {
	if l == 0 {
		return 1, &errorDetail{
			Nombre:      "emptystring",
			Descripcion: "Valor puede ser cero",
		}
	}
	return l, nil
}

func verifString(s string) (string, error) {
	if s == "" {
		return "", &errorDetail{
			Nombre:      "emptystring",
			Descripcion: "Valor puede estar vacío",
		}
	}
	return s, nil
}



func main() {

defer func(){
	fmt.Println("Fin de la ejercución")
	fmt.Println("Se detectaron varios errores en tiempo de ejecucion")
	fmt.Println("No han quedado archivos abiertos")
}()


}

