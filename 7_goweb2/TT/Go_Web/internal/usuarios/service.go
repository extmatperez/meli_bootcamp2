package internal

type service struct {
	respository Repository
}

type Service interface {
	GetAll() ([]Usuario, error)
	Store(nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error)
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
