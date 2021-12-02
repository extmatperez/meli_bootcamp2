// Repository pertenece al paquete internal (carpeta general a la que pertenece)
package internal

// Creamos la interface Service
type Service interface {
	Get_users() ([]Users, error)
	Post_users(first_name string, last_name string, email string, age int, height int, active bool, date string) (Users, error)
}
