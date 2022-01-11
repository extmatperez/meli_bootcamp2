package section

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/internal/domain"
)

type StubRepository struct {
	useGetAll bool
}

var sections string = `[
	{"id":1,"section_number": 54,"current_temperature": 33,"minimum_temperature": 100,"current_capacity": 10,"minimum_capacity": 25,"maximum_capacity": 13,"warehouse_id": 12,"product_type_id": 20},
	{"id":2,"section_number":3,"current_temperature":20,"minimum_temperature":14,"current_capacity":10,"minimum_capacity":10,"maximum_capacity":30,"warehouse_id":1,"id_product_type":2}]`

func (s *StubRepository) GetAll(ctx context.Context) ([]domain.Section, error) {
	var salida []domain.Section
	err := json.Unmarshal([]byte(sections), &salida)
	s.useGetAll = true
	return salida, err
}
func (s *StubRepository) Get(ctx context.Context, id int) (domain.Section, error) {
	var salida []domain.Section
	err := json.Unmarshal([]byte(sections), &salida)
	if err != nil {
		return domain.Section{}, err
	}
	s.useGetAll = true
	for _, section := range salida {
		if section.ID == id {
			return section, nil
		}
	}
	return domain.Section{}, errors.New("404")
}

func (s *StubRepository) Save(ctx context.Context, sec domain.Section) (int, error) {
	var salida []domain.Section
	err := json.Unmarshal([]byte(sections), &salida)
	if err != nil {
		return 0, err
	}
	s.useGetAll = true
	if s.Exists(ctx, sec.SectionNumber) {
		return -1, errors.New("409")
	}
	return sec.ID, nil
}

func (s *StubRepository) Update(ctx context.Context, sec domain.Section) error {
	var salida []domain.Section
	err := json.Unmarshal([]byte(sections), &salida)
	if err != nil {
		return err
	}
	s.useGetAll = true
	for _, section := range salida {
		if section.ID == sec.ID {
			return nil
		}
	}
	return errors.New("404")
}

func (s *StubRepository) Delete(ctx context.Context, id int) error {
	var salida []domain.Section
	if id == 0 {
		return errors.New("404")
	}
	err := json.Unmarshal([]byte(sections), &salida)
	if err != nil {
		return err
	}
	s.useGetAll = true
	for _, section := range salida {
		if section.ID == id {
			return nil
		}
	}
	return errors.New("404")
}

func (s *StubRepository) Exists(ctx context.Context, SectionNumber int) bool {
	var salida []domain.Section
	err := json.Unmarshal([]byte(sections), &salida)
	if err != nil {
		return false
	}
	for _, section := range salida {
		if section.SectionNumber == SectionNumber {
			return true
		}
	}
	return false
}
