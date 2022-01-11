package employee

// Archivo a modificar Server.
import (
	"context"
	"errors"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("employee not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Employee, error)
	Get(ctx context.Context, id int) (domain.Employee, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, empl domain.Employee) (int, error)
	Update(ctx context.Context, empl domain.Employee) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) GetAll(ctx context.Context) ([]domain.Employee, error) {
	return ser.repository.GetAll(ctx)
}

func (ser *service) Get(ctx context.Context, id int) (domain.Employee, error) {
	return ser.repository.Get(ctx, id)
}

func (ser *service) Exists(ctx context.Context, card_number_id string) bool {
	return ser.repository.Exists(ctx, card_number_id)
}

func (ser *service) Save(ctx context.Context, empl domain.Employee) (int, error) {
	idEmpl, err := ser.repository.Save(ctx, empl)
	if err != nil {
		return idEmpl, err
	}
	return idEmpl, nil
}

func (ser *service) Update(ctx context.Context, empl domain.Employee) error {
	return ser.repository.Update(ctx, empl)
}

func (ser *service) Delete(ctx context.Context, id int) error {
	return ser.repository.Delete(ctx, id)
}
