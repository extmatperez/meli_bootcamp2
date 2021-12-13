package main

import "fmt"

const (
	tarantula = "tarantula"
	hamster   = "hamster"
	perro     = "perro"
	gato      = "gato"
)

func main() {
	bicho := perro
	racion, err := animal(bicho)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	fmt.Printf("%s", racion())
}

func animal(animal string) (func() string, error) {
	switch animal {
	case perro:
		return racion_perro, nil
	case gato:
		return racion_gato, nil
	case hamster:
		return racion_hamster, nil
	case tarantula:
		return racion_tarantula, nil
	}
	return nil, fmt.Errorf("no tenemos ese animal en la tienda")
}

func racion_perro() string {
	return "El perro necesita 10kg de racion"
}

func racion_gato() string {
	return "El perro necesita 5kg de racion"
}

func racion_hamster() string {
	return "El perro necesita 250g de racion"
}

func racion_tarantula() string {
	return "El perro necesita 150g de racion"
}
