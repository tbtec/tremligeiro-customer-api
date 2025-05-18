package repository

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

type MockCustomerRepo struct {
	CreateFunc               func(ctx context.Context, customer *model.Customer) error
	FindByDocumentNumberFunc func(ctx context.Context, documentNumber string) (*model.Customer, error)
	FindOneFunc              func(ctx context.Context, id string) (*model.Customer, error)
}

func (m *MockCustomerRepo) Create(ctx context.Context, customer *model.Customer) error {
	return m.CreateFunc(ctx, customer)
}

func (m *MockCustomerRepo) FindByDocumentNumber(ctx context.Context, id string) (*model.Customer, error) {
	if m.FindByDocumentNumberFunc != nil {
		return m.FindByDocumentNumberFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockCustomerRepo) FindOne(ctx context.Context, id string) (*model.Customer, error) {
	if m.FindOneFunc != nil {
		return m.FindOneFunc(ctx, id)
	}
	return nil, nil
}
