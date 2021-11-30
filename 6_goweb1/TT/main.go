/*
Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder filtrar por todos los campos.
Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
Luego genera la lógica de filtrado de nuestro array.
Devolver por el endpoint el array filtrado.
*/

/*
Generar un nuevo endpoint que nos permita traer un solo resultado del array de la temática.
Utilizando path parameters el endpoint debería ser /temática/:id (recuerda que siempre tiene que ser en plural la temática).

Una vez recibido el id devuelve la posición correspondiente.
Genera una nueva ruta.
Genera un handler para la ruta creada.
Dentro del handler busca el item que necesitas.
Devuelve el item según el id.
Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.
*/

package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        string `json:"precio"`
	Stock         int    `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

func generarListaProductos() []Producto {

	data, _ := os.ReadFile("./productos.json")

	var lista []Producto

	json.Unmarshal(data, &lista)

	return lista
}

func filtrarProductosColor(c *gin.Context) {

	color := c.Param("color")

	productos := generarListaProductos()

	var productosFiltrados []*Producto

	for i, p := range productos {
		if p.Color == color {
			productosFiltrados = append(productosFiltrados, &productos[i])
		}
	}

	if len(productosFiltrados) != 0 {
		c.JSON(200, productosFiltrados)
	} else {
		c.JSON(404, "No se encontraron productos")
	}
}

func filtrarProductoId(c *gin.Context) {

	id := c.Param("id")

	productos := generarListaProductos()

	var productoEncontrado Producto

	for _, p := range productos {
		if strconv.Itoa(p.ID) == id {
			productoEncontrado = p
		}
	}

	if productoEncontrado.ID != 0 {
		c.JSON(200, productoEncontrado)
	} else {
		c.JSON(404, "No se encontró el producto")
	}
}

func GetAll(c *gin.Context) {

	lista := generarListaProductos()

	c.JSON(200, gin.H{
		"productos": lista,
	})
}

func main() {

	router := gin.Default()

	router.GET("/productos", GetAll)
	router.GET("/productos/:id", filtrarProductoId)
	router.GET("/filtrarColor/:color", filtrarProductosColor)

	router.Run("localhost:8080")
}
