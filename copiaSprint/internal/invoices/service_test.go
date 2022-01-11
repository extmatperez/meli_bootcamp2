package employee

import (
	"context"
	"testing"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestFindAllService(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	myEmployees, _ := service.GetAll(context.Background())

	assert.Equal(t, 2, len(myEmployees))
	assert.True(t, stubRepo.stubRepositoryFlag)
}

func TestFindByIDNonExistentService(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	_, err := service.Get(context.Background(), 3)

	assert.NotNil(t, err)
	assert.True(t, stubRepo.stubRepositoryFlag)
}

func TestFindByIDExistentService(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	myEmployee, err := service.Get(context.Background(), 1)

	assert.Equal(t, myEmployee.CardNumberID, "1234561")
	assert.Nil(t, err)
	assert.True(t, stubRepo.stubRepositoryFlag)
}

func TestCreateOKService(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	newEmployeeTest := domain.Employee{
		ID:           4,
		CardNumberID: "1234563",
		FirstName:    "Claudio",
		LastName:     "Castro",
		WarehouseID:  3,
	}

	idRecived, err := service.Save(context.Background(), newEmployeeTest)

	assert.Equal(t, newEmployeeTest.ID, idRecived)
	assert.Nil(t, err)
	assert.True(t, stubRepo.stubRepositoryFlag)
}

func TestCreateConflict(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	newEmployeeTest := domain.Employee{
		ID:           4,
		CardNumberID: "1234562",
		FirstName:    "Claudio",
		LastName:     "Castro",
		WarehouseID:  3,
	}
	errExists := service.Exists(context.Background(), newEmployeeTest.CardNumberID)

	assert.Equal(t, errExists, true)
	assert.True(t, stubRepo.stubRepositoryFlag)
}

func TestCreateAtributesNotFound(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	newEmployeeTest := domain.Employee{
		ID:           4,
		CardNumberID: "1234562",
		LastName:     "Castro",
		WarehouseID:  3,
	}
	_, errExists := service.Save(context.Background(), newEmployeeTest)

	assert.NotNil(t, errExists)
	assert.True(t, stubRepo.stubRepositoryFlag)
}

func TestUpdateExistent(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	newEmployeeTest := domain.Employee{
		ID:           2,
		CardNumberID: "1234562",
		FirstName:    "Claudio",
		LastName:     "Castro",
		WarehouseID:  3,
	}
	err := service.Update(context.Background(), newEmployeeTest)

	assert.Nil(t, err)
	assert.True(t, stubRepo.stubRepositoryFlag)
}

func TestUpdateNonExistent(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	newEmployeeTest := domain.Employee{
		ID:           4,
		CardNumberID: "1234562",
		FirstName:    "Claudio",
		LastName:     "Castro",
		WarehouseID:  3,
	}
	err := service.Update(context.Background(), newEmployeeTest)

	assert.NotNil(t, err)
	assert.True(t, stubRepo.stubRepositoryFlag)
}

func TestDeleteNonExistent(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(context.Background(), 22)

	assert.NotNil(t, err)
	assert.True(t, stubRepo.stubRepositoryFlag)
}

func TestDeleteOk(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(context.Background(), 2)

	assert.Nil(t, err)
	assert.True(t, stubRepo.stubRepositoryFlag)
}
