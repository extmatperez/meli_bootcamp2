// Es conveniente que nuestro codigo tenga un adecuado manejo de errores.
// "error" en Go es un tipo interface, por eso no se requieren estructuras de control rigidas.
// El unico metodo que tiene la interface error es un string que se llama Error().
// Por lo tanto, se pueden crear errores personalizados.

package main

import (
	"errors"
	"fmt"
)

// PAQUETE errors
// As(). Comprueba si un error es de tipo especifico.
// Is(). Verifica si un error tiene internamente a otro tipo de error, si lo contiene.
// Unwrap(). Devuelve el error subyacente o interno. Necesita un error y devuelve un error, osea, al error de dentro de otro error.

// Cuando se valida el error obtenido en una funcion siempre es recomendable ver que hacer en caso que err != nil, osea cuando hay error.
// El else por defecto de este camino es el camino feliz, pero es para definirlo.

type myError struct {
	status int    //
	msg    string //
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s - status: %d.", e.msg, e.status)
}

func devolver_error() error {
	return errors.New("Esto es un error creado por mi")
}

func devolver_custom_error() error {
	var errorcito myError
	errorcito.status = 500
	errorcito.msg = "Esto es un error creado por mi."
	// Aca pide la direccion del error creado para hacer referencia al mismo.
	// Aparte porque en la implementacion del metodo Error para myError uso el puntero en *myError.
	return &errorcito
}

// Esto nos va a servir para los errores personaliados cuando estemos con REST.
func myErrorTest(status int) (int, error) {
	if status >= 300 {
		return 400, &myError{
			status: 400,
			msg:    "Esto es una prueba de mi error.",
		}
	}
	return 200, nil
}

func main() {
	err1 := devolver_error()
	fmt.Println(err1)

	err2 := devolver_custom_error()
	fmt.Println(err2)

	num, err3 := myErrorTest(200)
	if err3 == nil {
		fmt.Println(num)
	} else {
		fmt.Println(err3)
	}

	// Tambien tenemos fmt.Errorf para formatear errores. Ver diapos.

	// Uso de As() del package Errors. err1 es del mismo tipo que err2, pero no err3.
	isMyError1 := errors.As(err1, &err2)
	fmt.Println(isMyError1)
	isMyError3 := errors.As(err3, &err2)
	fmt.Println(isMyError3)

	// Uso de Is() del package Errors. err1 esta contenido en err2, pero no err3 en err2.
	fmt.Println(errors.Is(err1, err2))
	fmt.Println(errors.Is(err3, err2))

	// Uso de Unwrap() del package Errors.
	fmt.Println(errors.Unwrap(err1))
	fmt.Println(errors.Unwrap(err2))
	fmt.Println(errors.Unwrap(err3))
}
