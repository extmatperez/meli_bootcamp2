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

////////////////////////////////////////

// Ejercicio 3 - Listar Entidad

// Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
// Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
// Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
// Genera un handler para el endpoint llamado “GetAll”.
// Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.
//	"github.com/gin-gonic/gin"

package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type Product []struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Price        string `json:"price"`
	Stock        string `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published"`
	CreationDate string `json:"creation_date"`
}

// var products []interface{}
// 		jsonFile, err := os.ReadFile("./products.json")
// 		if err == nil {
// 			json.Unmarshal(jsonFile, &products)
// 			c.JSON(http.StatusOK, products)
// 		}
// Diego Alejandro Parra Daza16:39
// 	archivo, _ := os.ReadFile("./products.json")
// 	var listProducts []Product
// 	json.Unmarshal(archivo, &listProducts)
// 	c.JSON(http.StatusOK, gin.H{
// 		"products": listProducts,
// 	})

func getAll(c *gin.Context) {
	var products []interface{}
	//file, _ := ioutil.ReadFile("./products.json")
	file, _ := os.ReadFile("./products.json")

	// if err := json.Unmarshal(file, &products); err != nil {
	// 	log.Fatal(err)
	// }
	json.Unmarshal([]byte(file), &products)

	c.JSON(200, gin.H{
		"data": products,
	})
}

// func getAll(c *gin.Context) {

// 	datos, _ := os.ReadFile("./productos.json")
// 	var lista []Productos
// 	json.Unmarshal(datos, &lista)

// 	c.JSON(http.StatusOK, gin.H{
// 		"productos": lista,
// 	})
// }

func main() {

	// file, err := ioutil.ReadFile("./products.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var as Product

	// jsonData, err := json.Marshal(file)

	// as = jsonData

	//ext := string(file)
	//fmt.Println(text)

	// Crea un router con gin
	router := gin.Default()
	// Captura la solicitud GET “/hello-world”
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Franco!",
		})
	})
	router.GET("/products", getAll)
	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()

}
