/*
El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes.

Los datos requeridos para registrar a un cliente son:
Legajo
Nombre y Apellido
DNI
Número de teléfono
Domicilio

Tarea 1: El número de legajo debe ser asignado o generado por separado y en forma previa a la carga de los restantes gastos.
Desarrolla e implementa una función para generar un ID que luego utilizarás para asignarlo como valor a “Legajo”.
Si por algún motivo esta función retorna valor “nil”, debe generar un panic que interrumpa la ejecución y aborte.

Tarea 2: Antes de registrar a un cliente, debes verificar si el mismo ya existe. Para ello, necesitas leer los datos de un archivo .txt.
En algún lugar de tu código, implementa la función para leer un archivo llamado “customers.txt” (como en el ejercicio anterior,
este archivo no existe, por lo que la función que intente leerlo devolverá un error). Debes manipular adecuadamente ese error
como hemos visto hasta aquí.

Ese error deberá:
1.- generar un panic;
2.- lanzar por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”, y continuar con la ejecución del programa normalmente.

Tarea 3: Luego de intentar verificar si el cliente a registrar ya existe, desarrolla una función para validar que todos los datos
a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores. Uno de los valores
retornados deberá ser de tipo error para el caso de que se ingrese por parámetro algún valor cero (recuerda los valores cero de cada
tipo de dato, ej: 0, “”, nil).

Tarea 4: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes:
“Fin de la ejecución”, “Se detectaron varios errores en tiempo de ejecución” y “No han quedado archivos abiertos” (en ese orden).
Utiliza defer para cumplir con este requerimiento.

Requerimientos generales:
Utiliza recover para recuperar el valor de los panics que puedan surgir (excepto en la tarea 1).
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error (por ejemplo las que intenten leer archivos).
Genera algún error, personalizandolo a tu gusto, utilizando alguna de las funciones que GO provee para ello (realiza también la validación
pertinente para el caso de error retornado).
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	Legajo    int    `json: "legajo"`
	Nombre    string `json: "nombre"`
	Apellido  string `json: "apellido"`
	Dni       int    `json: "dni"`
	Telefono  int    `json: "telefono"`
	Domicilio string `json: "domicilio"`
}

type myError struct {
	msg string
}

func (e myError) Error() string {
	panic(e.msg)
}

func generarLegajo() int {
	return rand.Int()
}

func leerArchivo() ([]Cliente, error) {

	data, err := os.ReadFile("./customers.txt")

	if err != nil {
		return nil, myError{"error: el archivo indicado no fue encontrado o está dañado"}
	} else {
		var pListaLeida []Cliente

		json.Unmarshal(data, &pListaLeida)

		return pListaLeida, nil
	}
}

func registrarCliente(legajo int, nuevoCliente Cliente) {

	listaLeida, err := leerArchivo()

	if err != nil {
		_ = err.Error()
	} else {

		for _, c := range listaLeida {
			if c.Legajo == nuevoCliente.Legajo {
				fmt.Printf("Ese usuario ya existe\n")
				return
			}
		}

		ok, err := validarDatos(nuevoCliente)

		if !ok {
			fmt.Printf("Error: %v\n", err)
			return
		}

		listaLeida = append(listaLeida, nuevoCliente)

		listaFormateada, _ := json.Marshal(listaLeida)

		err = os.WriteFile("../customers.txt", listaFormateada, 0644)

		if err != nil {
			fmt.Printf("No se pudo escribir el archivo\n")
		} else {
			fmt.Printf("Archivo escrito exitosamente\n")
		}
	}
}

func validarDatos(nuevoCliente Cliente) (bool, error) {

	if nuevoCliente.Legajo == 0 {
		return false, errors.New("el legajo no puede estar vacío")
	}

	if nuevoCliente.Nombre == "" {
		return false, errors.New("el nombre no puede estar vacío")
	}

	if nuevoCliente.Apellido == "" {
		return false, errors.New("el apellido no puede estar vacío")
	}

	if nuevoCliente.Dni == 0 {
		return false, errors.New("el dni no puede estar vacío")
	}

	if nuevoCliente.Telefono == 0 {
		return false, errors.New("el teléfono no puede estar vacío")
	}

	if nuevoCliente.Domicilio == "" {
		return false, errors.New("el domicilio no puede estar vacío")
	}

	return true, myError{}
}

func main() {

	defer func() {

		err := recover()

		fmt.Printf("%v\n", err)

		fmt.Printf("Fin de la ejecución\n")
		fmt.Printf("Se detectaron varios errores en tiempo de ejecición\n")
		fmt.Printf("No han quedado archivos abiertos\n")

	}()

	registrarCliente(generarLegajo(), Cliente{generarLegajo(), "Lucas", "Perez", 456966, 315986326, "General Paz 48"})

}
