package section

import (
	"context"
	"errors"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("section not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id int) (domain.Section, error)
	Exists(ctx context.Context, SectionNumber int) bool
	Save(ctx context.Context, s domain.Section) (int, error)
	Update(ctx context.Context, s domain.Section) error
	Delete(ctx context.Context, id int) error
}
type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (ser *service) Exists(ctx context.Context, SectionNumber int) bool {
	return ser.repository.Exists(ctx, SectionNumber)
}

func (ser *service) GetAll(ctx context.Context) ([]domain.Section, error) {
	sections, err := ser.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func (ser *service) Get(ctx context.Context, id int) (domain.Section, error) {
	sect, err := ser.repository.Get(ctx, id)
	if err != nil {
		return domain.Section{}, err
	}
	return sect, nil
}

func (ser *service) Save(ctx context.Context, s domain.Section) (int, error) {
	id, err := ser.repository.Save(ctx, s)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (ser *service) Update(ctx context.Context, sec domain.Section) error {
	return ser.repository.Update(ctx, sec)
}

func (ser *service) Delete(ctx context.Context, id int) error {
	return ser.repository.Delete(ctx, id)
}
