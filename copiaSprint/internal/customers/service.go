package buyer

import (
	"context"
	"errors"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("buyer not found")
)

type Service interface {
	GetAll(c context.Context) ([]domain.Buyer, error)
	Get(c context.Context, id int) (domain.Buyer, error)
	Save(c context.Context, buyerAux domain.Buyer) (int, error)
	Update(c context.Context, buyerAux domain.Buyer) (domain.Buyer, error)
	Delete(c context.Context, varID int) error
	Exist(c context.Context, CardNumberID string) bool
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) GetAll(c context.Context) ([]domain.Buyer, error) {
	response, err := ser.repository.GetAll(c)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (ser *service) Save(c context.Context, buyerAux domain.Buyer) (int, error) {
	response, err := ser.repository.Save(c, buyerAux)
	if err != nil {
		return 1, err
	}
	return response, err
}

func (ser *service) Get(c context.Context, id int) (domain.Buyer, error) {
	var buyerAux domain.Buyer
	response, err := ser.repository.Get(c, id)
	if err != nil {
		return buyerAux, err
	}
	return response, nil
}

func (ser *service) Update(c context.Context, buyerAux domain.Buyer) (domain.Buyer, error) {
	buyerToModificate, err := ser.repository.Get(c, buyerAux.ID)
	if err != nil {
		return buyerToModificate, err
	} else {
		buyerToModificate = buyerAux
		err := ser.repository.Update(c, buyerToModificate)
		if err != nil {
			return buyerToModificate, err
		}
		return buyerToModificate, nil
	}
}

func (ser *service) Delete(c context.Context, id int) error {
	err := ser.repository.Delete(c, id)
	if err != nil {
		return err
	}
	return nil
}

func (ser *service) Exist(c context.Context, card_number_id string) bool {
	return ser.repository.Exists(c, card_number_id)
}
