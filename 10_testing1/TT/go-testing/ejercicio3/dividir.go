package dividir

import "fmt"

// Cambiar el método para que no sólo retorne un entero sino también un error. Incorporar una validación en la que si el denominador es igual a 0,  retornar un error cuyo mensaje sea “El denominador no puede ser 0”. Diseñar un test unitario que valide el error cuando se invoca con 0 en el denominador.
// Dentro de la carpeta go-testing crear un archivo dividir.go con la función a probar.
// Dentro de la carpeta go-testing crear un archivo dividir_test.go con el test diseñado.

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("El denominador no puede ser 0")
	}
	return num / den, nil
}
