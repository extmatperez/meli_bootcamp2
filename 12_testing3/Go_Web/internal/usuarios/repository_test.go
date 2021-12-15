package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/11_testing2/Go_Web/pkg/store"
	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	ValidRead  bool
	ValidWrite bool
}

var users string = `
[{"id":1,"nombre":"Ida","apellido":"Tieman","email":"itieman0@npr.org","edad":82,"altura":187,"activo":true,"fecha_creacion":"06/15/2021"},
{"id":2,"nombre":"Law","apellido":"Lafee","email":"llafee1@barnesandnoble.com","edad":70,"altura":142,"activo":true,"fecha_creacion":"07/12/2021"}]

`

func (s *StubStore) Read(data interface{}) error {
	s.ValidRead = true
	return json.Unmarshal([]byte(users), &data)
}

func (s *StubStore) Write(data interface{}) error {
	s.ValidWrite = true
	return nil
}

func TestGetAll(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)
	//Act
	misUsuarios, _ := repo.GetAll()

	var expected []Usuario
	json.Unmarshal([]byte(users), &expected)

	//Assert
	assert.Equal(t, expected, misUsuarios)

}

func TestLastID(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)

	lastID := 2
	//Act
	ID, _ := repo.LastID()

	var expected []Usuario
	json.Unmarshal([]byte(users), &expected)

	//Assert
	assert.Equal(t, lastID, ID)
}
func TestUpdate(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)

	nombreEsperado := "pedro"

	//Act
	updateUser, err := repo.Update(1, nombreEsperado, "aponte", "email", 24, 123, true, "10-05-98")

	//Assert
	assert.Equal(t, nombreEsperado, updateUser.Nombre)
	assert.Nil(t, err)
}
func TestUpdateError(t *testing.T) {
	//Arrange
	store := StubStore{}
	repo := NewRepository(&store)

	nombreEsperado := "pedro"

	//Act
	_, err := repo.Update(5, nombreEsperado, "aponte", "email", 24, 123, true, "10-05-98")

	//Assert
	assert.Error(t, err)
}

func TestEditarNombreEdad(t *testing.T) {
	store := StubStore{false, false}

	repo := NewRepository(&store)

	nuevoNombre := "nuevo"
	nuevaEdad := 100

	actualizado, _ := repo.EditarNombreEdad(2, nuevoNombre, nuevaEdad)

	assert.Equal(t, Usuario{2, nuevoNombre, "Lafee", "llafee1@barnesandnoble.com", nuevaEdad, 142, true, "07/12/2021"}, actualizado)
	assert.True(t, store.ValidRead)
	assert.True(t, store.ValidWrite)
}

func TestGetAllRepositoryMock(t *testing.T) {
	dataByte := []byte(users)
	dbMock := store.Mock{Data: dataByte}
	var usuariosEsperados []Usuario
	json.Unmarshal(dataByte, &usuariosEsperados)

	storeStub := store.FileStore{Mock: &dbMock}

	repo := NewRepository(&storeStub)

	misUsuarios, _ := repo.GetAll()

	assert.Equal(t, usuariosEsperados, misUsuarios)
}

func TestGetLasIDRepositoryMock(t *testing.T) {
	dataByte := []byte(users)
	dbMock := store.Mock{Data: dataByte}
	var usuariosEsperados []Usuario
	json.Unmarshal(dataByte, &usuariosEsperados)

	storeStub := store.FileStore{Mock: &dbMock}

	repo := NewRepository(&storeStub)

	lastID, _ := repo.LastID()

	assert.Equal(t, usuariosEsperados[len(usuariosEsperados)-1].ID, lastID)
}
