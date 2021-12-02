package internal

type service struct {
	respository Repository
}

type Service interface {
	GetAll() ([]Usuario, error)
	Store(nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error)
	Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error)
	Delete(id int) error
}

func NewService(repository Repository) Service {
	return &service{respository: repository}
}

func (serv *service) GetAll() ([]Usuario, error) {
	usuarios, err := serv.respository.GetAll()
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (serv *service) Store(nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {
	lastID, err := serv.respository.LastID()
	if err != nil {
		return Usuario{}, err
	}
	newID := lastID + 1
	nuevoUsuario, err := serv.respository.Store(newID, nombre, apellido, email, edad, altura, activo, fecha)
	if err != nil {
		return Usuario{}, err
	}
	return nuevoUsuario, nil
}

func (serv *service) Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {
	updateUser, err := serv.respository.Update(id, nombre, apellido, email, edad, altura, activo, fecha)

	if err != nil {
		return Usuario{}, err
	}

	return updateUser, nil
}

func (serv *service) Delete(id int) error {
	err := serv.respository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
