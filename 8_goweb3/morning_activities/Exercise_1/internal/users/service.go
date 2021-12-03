// Repository pertenece al paquete internal (carpeta general a la que pertenece)
package internal

// Creamos la interface Service
type Service interface {
	Get_users() ([]Users, error)
	Post_users(first_name, last_name, email string, age, height int, active bool, date string) (Users, error)
	Update_users(id int, first_name, last_name, email string, age, height int, active bool, date string) (Users, error)
	Update_users_first_name(id int, first_name string) (Users, error)
}

// Agregamos la struct service
type service struct {
	repository Repository
}

// Creamos el New_service al cual le pasamos el repositorio
func New_service(repository Repository) Service {
	return &service{repository: repository}
}

// Creamos los métodos del New_service
// Llamamos al repository para que proceda a llamar a todos los usuarios
func (ser *service) Get_users() ([]Users, error) {
	users, err := ser.repository.Get_users()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Llamamos al repository para que proceda a agregar a un usuario
func (ser *service) Post_users(first_name, last_name, email string, age, height int, active bool, date string) (Users, error) {
	last_id, err := ser.repository.Last_id()
	if err != nil {
		return Users{}, err
	}
	user, err := ser.repository.Post_users(last_id+1, first_name, last_name, email, age, height, active, date)
	if err != nil {
		return Users{}, err
	}
	return user, nil
}

// Llamamos al repository para que proceda a actualizar al usuario
func (ser *service) Update_users(id int, first_name, last_name, email string, age, height int, active bool, date string) (Users, error) {
	return ser.repository.Update_users(id, first_name, last_name, email, age, height, active, date)
}

// Llamamos al repository para que proceda a actualizar al nombre del usuario
func (ser *service) Update_users_first_name(id int, first_name string) (Users, error) {
	return ser.repository.Update_users_first_name(id, first_name)
}
