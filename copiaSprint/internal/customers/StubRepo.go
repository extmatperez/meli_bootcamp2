package buyer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/internal/domain"
)

type StubRepository struct {
	useGetAll bool
}

var buyerServiceTest string = `[
    {"id":1,"card_number_id": "34203600","first_name": "Damian","last_name":"Zamora"},
    {"id":2,"card_number_id": "33200200","first_name": "Daniel","last_name":"Lopez"}]`

func (s *StubRepository) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	var salida []domain.Buyer
	err := json.Unmarshal([]byte(buyerServiceTest), &salida)
	s.useGetAll = true
	return salida, err
}

func (s *StubRepository) Get(ctx context.Context, id int) (domain.Buyer, error) {
	var salida []domain.Buyer
	json.Unmarshal([]byte(buyerServiceTest), &salida)
	fmt.Print(len(salida))
	for i := 0; i < len(salida); i++ {
		if id == salida[i].ID {
			s.useGetAll = true
			return salida[i], nil
		}
	}
	s.useGetAll = true
	return salida[0], fmt.Errorf("No se encontro ID")

}

func (s *StubRepository) Save(ctx context.Context, buyerAux domain.Buyer) (int, error) {
	var baseDeDatos []domain.Buyer
	json.Unmarshal([]byte(buyerServiceTest), &baseDeDatos)
	for i := 0; i < len(baseDeDatos); i++ {
		if buyerAux.CardNumberID == baseDeDatos[i].CardNumberID {
			return 1, errors.New("409")
		}
	}
	baseDeDatos = append(baseDeDatos, buyerAux)
	return len(baseDeDatos), nil

}
func (s *StubRepository) Update(ctx context.Context, buyerAux domain.Buyer) error {
	var baseDeDatos []domain.Buyer
	json.Unmarshal([]byte(buyerServiceTest), &baseDeDatos)
	for i := 0; i < len(baseDeDatos); i++ {
		if buyerAux.ID == baseDeDatos[i].ID {
			return nil
		}
	}
	return errors.New("404")
}

func (s *StubRepository) Delete(ctx context.Context, id int) error {
	var salida []domain.Buyer
	json.Unmarshal([]byte(buyerServiceTest), &salida)
	for i := 0; i < len(salida); i++ {
		if id == salida[i].ID {
			return nil
		}
	}
	return fmt.Errorf("No se pudo eliminar")
}
func (s *StubRepository) Exists(ctx context.Context, CardNumberID string) bool {
	var buyers []domain.Buyer
	err := json.Unmarshal([]byte(buyerServiceTest), &buyers)
	if err != nil {
		return false
	}
	for _, buyer := range buyers {
		if buyer.CardNumberID == CardNumberID {
			return true
		}
	}
	return false
}
