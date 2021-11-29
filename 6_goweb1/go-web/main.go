// Ejercicio 1 - Estructura un JSON
// Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.
// Los productos varían por id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no), fecha de creación.
// Los usuarios varían por id, nombre, apellido, email, edad, altura, activo (si-no), fecha de creación.
// Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor (string), fecha de transacción.
// Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el tema elegido, ej: products.json.
// Dentro del mismo escribí un JSON que permita tener un array de productos, usuarios o transacciones con todas sus variantes.
/////////////////////////////////////////

// Ejercicio 2 - Hola {nombre}

// Crea dentro de la carpeta go-web un archivo llamado main.go
// Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
// Pegale al endpoint para corroborar que la respuesta sea la correcta.

package main

import "github.com/gin-gonic/gin"

func main() {
	// Crea un router con gin
	router := gin.Default()
	// Captura la solicitud GET “/hello-world”
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Franco!",
		})
	})
	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
