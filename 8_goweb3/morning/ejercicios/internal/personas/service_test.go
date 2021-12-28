package internal

import (
	"context"
	"encoding/json"
	"testing"
	"time"
	//"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/internal/models"
	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/internal/models"
	"github.com/stretchr/testify/assert"
)

type stubRepository struct {
	useGetAll bool
}

var perso2 string = `[
{
	"id": 1,
	"nombre": "donald",
	"apellido": "trump",
	"edad": 27
   },
   {
	"id": 2,
	"nombre": "jair",
	"apellido": "bolsonaro",
	"edad": 55
   }
]`

func (s *stubRepository) GetAll() ([]Persona, error) {
	var salida []Persona
	err := json.Unmarshal([]byte(perso2), &salida)
	s.useGetAll = true

	return salida, err
}

func (s *stubRepository) Store(id int, nombre string, apellido string, edad int) (Persona, error) {
	return Persona{}, nil
}
func (s *stubRepository) Update(id int, nombre string, apellido string, edad int) (Persona, error) {
	return Persona{}, nil
}
func (s *stubRepository) UpdateNombre(id int, nombre string) (Persona, error) {
	return Persona{}, nil
}
func (s *stubRepository) Delete(id int) error {
	return nil
}
func (s *stubRepository) LastId() (int, error) {
	return 0, nil
}

func (s *stubRepository) Average() (float64, error) {
	return 0.0, nil
}

func TestGetAll(t *testing.T) {
	stubRepo := stubRepository{false}
	service := NewService(&stubRepo)

	misPersonas, _ := service.GetAll()

	assert.Equal(t, 2, len(misPersonas))
	assert.True(t, stubRepo.useGetAll)
}

func TestDelete(t *testing.T) {
	stubRepo := stubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(1)

	assert.Nil(t, err)
}

func TestUpdateServiceSQL(t *testing.T) {
	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	newPersona := models.Persona{
		ID: 1,
		Nombre: "JC",
		Apellido: "Rossi",
		Edad: 21,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	personaUpdated, err := service.UpdateWithContext(ctx, newPersona)

	assert.Equal(t, personaUpdated.Nombre, newPersona.Nombre)
	assert.Nil(t, err)
}

/* func TestGetByNameServiceSQL(t *testing.T) {
	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	personaGet := service.GetByName("JC")

	fmt.Println(personaGet)

	assert.Equal(t, personaGet[0].Nombre, "JC")
} */

/* func TestDeleteServiceSQL(t *testing.T) {
	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	err := service.Delete(4)

	assert.Nil(t, err)
} */

/* func TestGetAllSQL(t *testing.T) {
	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	arr, err := service.GetAll()

	assert.Len(t, arr, 2)
	assert.Nil(t, err)
} */

/* func TestGetFullDataJoinedServiceSQL(t *testing.T) {
	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	arr, err := service.GetFullDataJoined()

	assert.Len(t, arr, 2)
	assert.Nil(t, err)

	fmt.Println(arr)
} */

/* func TestStoreServiceSQL(t *testing.T) {
	personaNueva := models.Persona{
		Nombre: "Donald",
		Apellido: "Trump",
		Edad: 68,
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	personaCreada, err := service.Store(personaNueva)

	//2do argumento: Lo que espero, 3er argumento: Lo que viene de la consulta
	assert.Equal(t, personaNueva.Nombre, personaCreada.Nombre)
	assert.Nil(t, err)
} */

/* func TestGetOneServiceSQL(t *testing.T) {
	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	personaGet, err := service.GetOne(5)

	fmt.Println(personaGet)
	fmt.Println(err)

	assert.Equal(t, personaGet.Nombre, "Donald")
	//assert.Nil(t, err)
} */

/* func TestGetOneWithContextServiceSQL(t *testing.T) {
	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	personaGet, err := service.GetOneWithContext(ctx, 5)
	assert.Nil(t, err)
	fmt.Println(personaGet)
} */