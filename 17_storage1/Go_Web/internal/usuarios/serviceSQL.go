package internal

import "github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/17_storage1/Go_Web/internal/models"

type ServiceSQL interface {
	Store(nombre, apellido, email string, edad, altura int, activo bool, fecha string) (models.Usuario, error)
	GetOne(id int) models.Usuario
	Update(persona models.Usuario) (models.Usuario, error)
	GetByName(nombre string) ([]models.Usuario, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (serv *serviceSQL) Store(nombre, apellido, email string, edad, altura int, activo bool, fecha string) (models.Usuario, error) {
	newUsuario := models.Usuario{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fecha}
	userCreated, err := serv.repository.Store(newUsuario)
	if err != nil {
		return models.Usuario{}, nil
	}
	return userCreated, nil
}

func (ser *serviceSQL) GetOne(id int) models.Usuario {
	return ser.repository.GetOne(id)
}

func (ser *serviceSQL) Update(persona models.Usuario) (models.Usuario, error) {
	return ser.repository.Update(persona)
}

func (ser *serviceSQL) GetByName(nombre string) ([]models.Usuario, error) {
	usuario, err := ser.repository.GetByName(nombre)
	if err != nil {
		return nil, err
	}

	return usuario, nil
}
