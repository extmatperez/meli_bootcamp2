package product

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/internal/domain"
)

type StubRepository struct {
	useGetAll bool
}

var prod = `[
	{	"Id":1,
		"description":"articulo test",
		"expiration_rate":3,
		"freezing_rate":5,
		"height":5,
		"length":66.2,
		"netweight":66.3,
		"product_code":"a41",
		"recommended_freezing_temperature":13.3,
		"width":2.5,
		"product_type_id":11,
		"seller_id":0
		},
	{
		"Id":2,
		"description":"articulo test2",
		"expiration_rate":4,
		"freezing_rate":6,
		"height":6,
		"length":67.2,
		"netweight":67.3,
		"product_code":"a42",
		"recommended_freezing_temperature":14.3,
		"width":3.5,
		"product_type_id":12,
		"seller_id":0
	}
]`

func (s *StubRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var salida []domain.Product
	err := json.Unmarshal([]byte(prod), &salida)
	s.useGetAll = true
	return salida, err
}
func (s *StubRepository) Get(ctx context.Context, id int) (domain.Product, error) {
	var salida []domain.Product
	err := json.Unmarshal([]byte(prod), &salida)
	s.useGetAll = true
	if id > len(salida) {
		return domain.Product{}, errors.New("Not exist")
	}
	return salida[id-1], err
}

func (s *StubRepository) Save(ctx context.Context, p domain.Product) (int, error) {
	var salida []domain.Product
	err := json.Unmarshal([]byte(prod), &salida)
	if err != nil {
		return 0, err
	}
	s.useGetAll = true
	if s.Exists(ctx, p.ProductCode) {
		return 0, errors.New("Product exist")
	}
	return p.ID, nil

}
func (s *StubRepository) Update(ctx context.Context, p domain.Product) error {
	var salida []domain.Product
	err := json.Unmarshal([]byte(prod), &salida)
	if err != nil {
		return err
	}
	s.useGetAll = true
	for _, producto := range salida {
		if producto.ID == p.ID {
			return nil
		}
	}
	return errors.New("noN exist")
}
func (s *StubRepository) Delete(ctx context.Context, id int) error {
	var salida []domain.Product
	err := json.Unmarshal([]byte(prod), &salida)
	if err != nil {
		return err
	}
	s.useGetAll = true
	for _, producto := range salida {
		if producto.ID == id {
			return nil
		}
	}
	return errors.New("Non exist")
}
func (s *StubRepository) Exists(ctx context.Context, productCode string) bool {
	var salida []domain.Product
	err := json.Unmarshal([]byte(prod), &salida)
	if err != nil {
		return false
	}
	for _, product := range salida {
		if product.ProductCode == productCode {
			return true
		}
	}
	return false
}
