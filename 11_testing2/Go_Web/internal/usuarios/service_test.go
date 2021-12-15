package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/11_testing2/Go_Web/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
	usedGetAll bool
}

var userss string = `
[{"id":1,"nombre":"Ida","apellido":"Tieman","email":"itieman0@npr.org","edad":82,"altura":187,"activo":true,"fecha_creacion":"06/15/2021"},
{"id":2,"nombre":"Law","apellido":"Lafee","email":"llafee1@barnesandnoble.com","edad":70,"altura":142,"activo":true,"fecha_creacion":"07/12/2021"}]

`

func (s *StubRepository) GetAll() ([]Usuario, error) {
	var salida []Usuario
	err := json.Unmarshal([]byte(userss), &salida)
	s.usedGetAll = true
	return salida, err
}

func (s *StubRepository) Store(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {
	return Usuario{}, nil
}

func (s *StubRepository) Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (Usuario, error) {
	return Usuario{}, nil
}

func (s *StubRepository) Delete(id int) error {
	return nil
}

func (s *StubRepository) EditarNombreEdad(id int, nombre string, edad int) (Usuario, error) {
	return Usuario{}, nil
}

func (s *StubRepository) LastID() (int, error) {
	return 0, nil
}

func TestGetAllService(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	misUsuarios, _ := service.GetAll()
	assert.Equal(t, 2, len(misUsuarios))
	assert.True(t, stubRepo.usedGetAll)
}

func TestGetAllServiceMock(t *testing.T) {
	dataByte := []byte(userss)
	var usuariosEsperados []Usuario
	json.Unmarshal(dataByte, &usuariosEsperados)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	misUsuarios, _ := service.GetAll()
	assert.Equal(t, usuariosEsperados, misUsuarios)
}
func TestGetAllServiceMockError(t *testing.T) {
	// dataByte := []byte(userss)
	// var usuariosEsperados []Usuario
	// json.Unmarshal(dataByte, &usuariosEsperados)
	errorEsperado := errors.New("No hay datos en el mock")

	dbMock := store.Mock{Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	misUsuarios, errorRecibido := service.GetAll()
	assert.Equal(t, errorEsperado, errorRecibido)
	assert.Nil(t, misUsuarios)
}

func TestStoreServiceMock(t *testing.T) {
	// Arrange
	usuarioNuevo := Usuario{3, "Juan", "Perez", "correo", 20, 180, true, "fecha"}

	dbMock := store.Mock{Data: []byte(`[]`)}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	usuarioCreado, _ := service.Store(usuarioNuevo.Nombre, usuarioNuevo.Apellido, usuarioNuevo.Email, usuarioNuevo.Edad, usuarioNuevo.Altura, usuarioNuevo.Activo, usuarioNuevo.FechaCreacion)
	assert.Equal(t, usuarioNuevo.Nombre, usuarioCreado.Nombre)
	assert.Equal(t, usuarioNuevo.Altura, usuarioCreado.Altura)
	//assert.Nil(t, misUsuarios)
}

func TestStoreServiceMockError(t *testing.T) {
	// Arrange
	usuarioNuevo := Usuario{3, "Juan", "Perez", "correo", 20, 180, true, "fecha"}

	errorEsperado := errors.New("No hay datos en el mock")

	dbMock := store.Mock{Data: []byte(`[]`), Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	usuarioCreado, err := service.Store(usuarioNuevo.Nombre, usuarioNuevo.Apellido, usuarioNuevo.Email, usuarioNuevo.Edad, usuarioNuevo.Altura, usuarioNuevo.Activo, usuarioNuevo.FechaCreacion)
	assert.Equal(t, errorEsperado, err)
	assert.Equal(t, "", usuarioCreado.Nombre)
}

func TestUpdateServiceMock(t *testing.T) {
	// Arrange
	usuarioUp := Usuario{2, "Juan", "Perez", "correo", 20, 180, true, "fecha"}

	dataByte := []byte(userss)
	var usuariosEsperados []Usuario
	json.Unmarshal(dataByte, &usuariosEsperados)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	usuarioActualizado, _ := service.Update(usuarioUp.ID, usuarioUp.Nombre, usuarioUp.Apellido, usuarioUp.Email, usuarioUp.Edad, usuarioUp.Altura, usuarioUp.Activo, usuarioUp.FechaCreacion)
	assert.Equal(t, usuarioUp, usuarioActualizado)
	assert.True(t, dbMock.ValidRead)
}

func TestDeleteServiceMock(t *testing.T) {
	// Arrange

	dataByte := []byte(userss)
	var usuariosEsperados []Usuario
	json.Unmarshal(dataByte, &usuariosEsperados)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	err := service.Delete(1)
	assert.Nil(t, err)
}
func TestDeleteServiceMockError(t *testing.T) {
	// Arrange

	dataByte := []byte(userss)
	var usuariosEsperados []Usuario
	json.Unmarshal(dataByte, &usuariosEsperados)
	errorEsperado := errors.New("No hay datos en el mock")

	dbMock := store.Mock{Data: dataByte, Err: errorEsperado}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	err := service.Delete(5)
	assert.Equal(t, errorEsperado, err)
}
func TestDeleteService(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(1)

	assert.Nil(t, err)
}
