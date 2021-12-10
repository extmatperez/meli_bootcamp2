/*
Cambiar el método para que no sólo retorne un entero sino también un error. Incorporar una validación en la que si el denominador es igual a 0,  retornar un error cuyo mensaje sea “El denominador no puede ser 0”. Diseñar un test unitario que valide el error cuando se invoca con 0 en el denominador.
Dentro de la carpeta go-testing crear un archivo dividir.go con la función a probar.
Dentro de la carpeta go-testing crear un archivo dividir_test.go con el test diseñado.
*/

package dividir

import "errors"

func Dividir(num1 int, num2 int) (int, error) {

	if num2 == 0 {
		return 0, errors.New("el denominador no puede ser 0")
	}

	return (num1 / num2), nil
}
