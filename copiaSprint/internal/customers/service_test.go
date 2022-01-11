package buyer

import (
	"context"
	"testing"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/assert"
)

//funciones Test
func Test_create_ok(t *testing.T) {
	newBuyer := domain.Buyer{
		CardNumberID: "34203601",
		FirstName:    "Negro",
		LastName:     "Last",
	}
	stubRepo := StubRepository{false}
	serviceStub := NewService(&stubRepo)
	idNueva, _ := serviceStub.Save(context.Background(), newBuyer)
	assert.Equal(t, 4, idNueva)

}

func Test_create_conflict(t *testing.T) {
	newBuyer := domain.Buyer{
		CardNumberID: "34203600",
		FirstName:    "Negro",
		LastName:     "Last",
	}

	stubRepo := StubRepository{false}
	serviceStub := NewService(&stubRepo)
	id, _ := serviceStub.Save(context.Background(), newBuyer)
	assert.Equal(t, 1, id)
}

func Test_find_all(t *testing.T) {

	stubRepo := StubRepository{false}
	serviceStub := NewService(&stubRepo)
	buyers, _ := serviceStub.GetAll(context.Background())
	assert.Equal(t, 3, len(buyers))
	assert.True(t, stubRepo.useGetAll)

}

func Test_find_by_id_existent(t *testing.T) {
	stubRepo := StubRepository{false}
	serviceStub := NewService(&stubRepo)
	buyer, _ := serviceStub.Get(context.Background(), 1)
	assert.Equal(t, 1, buyer.ID)
	assert.True(t, stubRepo.useGetAll)
}

func Test_find_by_id_non_existent(t *testing.T) {
	stubRepo := StubRepository{false}
	serviceStub := NewService(&stubRepo)
	_, err := serviceStub.Get(context.Background(), 5)
	assert.NotNil(t, err)
	assert.True(t, stubRepo.useGetAll)
}

func Test_update_ok(t *testing.T) {
	newBuyer := domain.Buyer{
		CardNumberID: "34203600",
		FirstName:    "Lopez",
		LastName:     "Nadia",
	}
	idACambiar := 2
	newBuyer.ID = idACambiar
	stubRepo := StubRepository{false}
	serviceStub := NewService(&stubRepo)
	buyerActualizado, _ := serviceStub.Update(context.Background(), newBuyer)
	assert.Equal(t, newBuyer.FirstName, buyerActualizado.FirstName)
}

func Test_update_non_existent(t *testing.T) {
	newBuyer := domain.Buyer{
		CardNumberID: "34203600",
		FirstName:    "Lopez",
		LastName:     "Nadia",
	}
	idACambiar := 8
	newBuyer.ID = idACambiar
	stubRepo := StubRepository{false}
	serviceStub := NewService(&stubRepo)
	_, err := serviceStub.Update(context.Background(), newBuyer)
	assert.NotNil(t, err)
}

func Test_delete_ok(t *testing.T) {
	stubRepo := StubRepository{false}
	serviceStub := NewService(&stubRepo)
	err := serviceStub.Delete(context.Background(), 1)
	assert.Nil(t, err)
}

func Test_delete_non_existent(t *testing.T) {
	stubRepo := StubRepository{false}
	serviceStub := NewService(&stubRepo)
	err := serviceStub.Delete(context.Background(), 8)
	assert.NotNil(t, err)
}
