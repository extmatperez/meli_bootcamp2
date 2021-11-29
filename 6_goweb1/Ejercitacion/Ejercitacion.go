// API y protocolo HTTP.
// 6 claves de la arquitectura web segun Roy Fielding.
// 1. Cliente-servidor.
// 2. Interfaz uniforme. Todos los nombres del servidor tienen un nombre en forme de URL o hipervinculo, un endpoint
// 3. Un sistema de capas. Los intermediarios en la red como proxies o gateways se implementan de manera transparente entre un cliente y un servidor.
// 4. Cache. Es un recurso utilizado por browsers y servidores para reducir el ancho de banda utilizado al cargar una pagina web. Reduce el costo total de la web.
// 5. Sin estado, o Stateless. Las interacciones entre el cliente y el servidor deben ser tratadas como nuevas y de forma independiente sin guardar estado.
// 6. Codigo a demanda. El cliente debe ser capaz de comprender y ejecutar el codigo que descarga bajo demanda del servidor.

// Siempre hay que descargarse el paquete json.
// Marshal. Toma los datos y los transforma en json
// Unmarshal. Toma los json y los transforma en datos tipados si se los asigno a una estructura que concuerde con lo devuelto.

package main

import (
	"fmt"
	"net/http"
)

func funcionSaludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo!")
}

func main() {
	// Este es el metodo con HTTP, que es justamente el que no se usa mas.
	http.HandleFunc("/hola", funcionSaludar)
	http.ListenAndServe(":8080", nil)
}
