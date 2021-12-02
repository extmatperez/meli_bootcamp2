// Repository pertenece al paquete internal (carpeta general a la que pertenece)
package internal

// Estructura de los datos que voy a manipular
type Users struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	Active    bool   `json:"active"`
	Date      string `json:"date"`
}

// Mientras no tengamos base de datos necesitamos saber donde almacenar los datos
var users []Users
var last_id int

// Definimos métodos a utilizar en nuestro repo
type Repository interface {
	get_users() ([]Users, error)
	validate_fields(user_id Users) (string, error)
	post_users() (Users, error)
}

// Agregamos la estructura repository donde vamos a tener implementados los métodos de la interface
type repository struct{}
