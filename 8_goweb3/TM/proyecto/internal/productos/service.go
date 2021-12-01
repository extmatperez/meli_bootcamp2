/*
Se solicita implementar una funcionalidad que modifique completamente una entidad. Para
lograrlo, es necesario, seguir los siguientes pasos:
	1. Generar un método PUT para modificar la entidad completa
	2. Desde el Path enviar el ID de la entidad que se modificará
	3. En caso de no existir, retornar un error 404
	4. Realizar todas las validaciones (todos los campos son requeridos)
*/

/*
Es necesario implementar una funcionalidad para eliminar una entidad. Para lograrlo, es
necesario, seguir los siguientes pasos:
	1. Generar un método DELETE para eliminar la entidad en base al ID
	2. En caso de no existir, retornar un error 404
*/

/*
Se requiere implementar una funcionalidad que modifique la entidad parcialmente, solo se
deben modificar 2 campos:
	- Si se seleccionó Productos, los campos nombre y precio.
	- Si se seleccionó Usuarios, los campos apellido y edad.
	- Si se seleccionó Transacciones, los campos código de transacción y monto.

Para lograrlo, es necesario, seguir los siguientes pasos:
	1. Generar un método PATCH para modificar la entidad parcialmente, modificando solo 2
	campo (a elección)
	2. Desde el Path enviar el ID de la entidad que se modificara
	3. En caso de no existir, retornar un error 404
	4. Realizar las validaciones de los 2 campos a enviar
*/

package internal

type Service interface {
	GetAll() ([]Producto, error)
	Store(nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Update(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Delete(id int) error
	UpdateNombrePrecio(id int, nombre string, precio string) (Producto, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]Producto, error) {

	productos, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return productos, nil
}

func (s *service) Store(nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	lastId, err := s.repository.LastId()

	if err != nil {
		return Producto{}, err
	}

	nuevoProducto, err := s.repository.Store(lastId+1, nombre, color, precio, stock, codigo, publicado, fechaCreacion)

	if err != nil {
		return Producto{}, err
	}

	return nuevoProducto, nil
}

func (s *service) Update(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {
	return s.repository.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateNombrePrecio(id int, nombre string, precio string) (Producto, error) {
	return s.repository.UpdateNombrePrecio(id, nombre, precio)
}
