package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("%vgrs\n", alimento("perro", 4))
	fmt.Printf("%vgrs\n", alimento("gato", 3))
	fmt.Printf("%vgrs\n", alimento("hamster", 2))
	fmt.Printf("%vgrs\n", alimento("tarantula", 5))
}

func alimento(mascota string, cantidad int) string {
	alimentoTotal, err := calcularAlimento(mascota, cantidad)
	if err != nil {
		return string(err.Error())
	}
	return fmt.Sprintf("%.3f", alimentoTotal)
}

func calcularAlimento(mascota string, cantidad int) (float64, error) {
	switch mascota {
	case "perro":
		return comidaPerro(cantidad), nil
	case "gato":
		return comidaGato(cantidad), nil
	case "tarantula":
		return comidaTar(cantidad), nil
	case "hamster":
		return comidaHam(cantidad), nil
	default:
		return 0, errors.New("No existe ese animal en la tienda")
	}
}

func comidaPerro(cantidad int) float64 {
	return float64(cantidad) * 10000
}

func comidaGato(cantidad int) float64 {
	return float64(cantidad) * 5000
}

func comidaHam(cantidad int) float64 {
	return float64(cantidad) * 250
}

func comidaTar(cantidad int) float64 {
	return float64(cantidad) * 150
}
