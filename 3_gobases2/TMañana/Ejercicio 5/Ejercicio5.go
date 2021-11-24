package main

import "fmt"
import "errors"


const (
	perro = "perro"
	gato = "gato"
	tarantula = "tarantula"
	hamster = "hamster"
 )
 
 
func dog(cantidad int) string {
	return "Debes comprar "+ fmt.Sprint(cantidad * 10) + "kg"
}


func cat(cantidad int) string {
	return "Debes comprar "+ fmt.Sprint(cantidad * 5) + "kg"
}

func hams(cantidad int) string {
	return "Debes comprar "+ fmt.Sprint(cantidad * 250) + "g"
}

func spider(cantidad int) string {
	return "Debes comprar "+ fmt.Sprint(cantidad * 150) + "g"
}



func Animal(nombre string) (func(cantidad int)string, error){
	switch nombre {
	case perro:
		return dog, nil
	case gato:
		return cat, nil
	case hamster:
		return hams, nil
	case tarantula:
		return spider, nil
	}
	return nil, errors.New("Operacion invalida")
}
	



func main() {
	
}