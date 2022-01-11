package section

import (
	"context"
	"testing"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_find_all(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	misSections, _ := service.GetAll(context.Background())

	assert.Equal(t, 2, len(misSections))
	assert.True(t, stubRepo.useGetAll)
}

func Test_find_by_id_existent(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	section, err := service.Get(context.Background(), 2)

	assert.Equal(t, 3, section.SectionNumber)
	assert.True(t, stubRepo.useGetAll)
	assert.Nil(t, err)
}
func Test_find_by_id_non_existent(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	section, err := service.Get(context.Background(), 3)
	assert.NotNil(t, err)
	assert.True(t, stubRepo.useGetAll)
	assert.Equal(t, 0, section.SectionNumber)
}

func Test_create_ok(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)
	sectionNuevo := domain.Section{
		ID:                 3,
		SectionNumber:      23,
		CurrentTemperature: 33,
		MinimumTemperature: 22,
		CurrentCapacity:    41,
		MinimumCapacity:    22,
		MaximumCapacity:    33,
		WarehouseID:        1,
		ProductTypeID:      2,
	}

	sectionID, err := service.Save(context.Background(), sectionNuevo)

	assert.Equal(t, sectionNuevo.ID, sectionID)
	assert.Nil(t, err)
	assert.True(t, stubRepo.useGetAll)
}
func Test_create_conflict(t *testing.T) {

	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)
	sectionNuevo := domain.Section{
		ID:                 4,
		SectionNumber:      54,
		CurrentTemperature: 33,
		MinimumTemperature: 22,
		CurrentCapacity:    41,
		MinimumCapacity:    22,
		MaximumCapacity:    33,
		WarehouseID:        1,
		ProductTypeID:      2,
	}

	sectionID, err := service.Save(context.Background(), sectionNuevo)

	assert.Equal(t, -1, sectionID)
	assert.NotNil(t, err)
	assert.True(t, stubRepo.useGetAll)
}

func Test_update_existent(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)
	sectionNuevo := domain.Section{
		ID:                 2,
		SectionNumber:      54,
		CurrentTemperature: 33,
		MinimumTemperature: 22,
		CurrentCapacity:    41,
	}

	err := service.Update(context.Background(), sectionNuevo)

	assert.Nil(t, err)
	assert.True(t, stubRepo.useGetAll)
}

func Test_update_non_existent(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)
	sectionNuevo := domain.Section{
		ID:                 20,
		SectionNumber:      54,
		CurrentTemperature: 33,
		MinimumTemperature: 22,
		CurrentCapacity:    41,
	}

	err := service.Update(context.Background(), sectionNuevo)

	assert.NotNil(t, err)
	assert.True(t, stubRepo.useGetAll)
}

func Test_delete_non_existent(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(context.Background(), 5)

	assert.NotNil(t, err)
	assert.True(t, stubRepo.useGetAll)
}

func Test_delete_ok(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)

	err := service.Delete(context.Background(), 1)

	assert.Nil(t, err)
	assert.True(t, stubRepo.useGetAll)
}

func Test_exists(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)
	boolean := service.Exists(context.Background(), 54)
	assert.True(t, boolean)
}
func Test_non_exists(t *testing.T) {
	stubRepo := StubRepository{false}
	service := NewService(&stubRepo)
	boolean := service.Exists(context.Background(), 24)
	assert.False(t, boolean)
}
