package repository

import (
	"context"
	"errors"

	"github.com/stretchr/testify/assert"
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

// Mock compatível com a interface ICustomerRepository
type MockCustomerRepoInterface struct{}

func (m *MockCustomerRepoInterface) Create(ctx context.Context, p *model.Customer) error { return nil }
func (m *MockCustomerRepoInterface) FindOne(ctx context.Context, id string) (*model.Customer, error) {
	return nil, nil
}
func (m *MockCustomerRepoInterface) FindByDocumentNumber(ctx context.Context, id string) (*model.Customer, error) {
	return nil, nil
}

// Mock customer repo that returns error
type MockCustomerRepoError struct{}

func (m *MockCustomerRepoError) Create(ctx context.Context, p *model.Customer) error {
	return assert.AnError
}

func (m *MockCustomerRepoError) FindOne(ctx context.Context, id string) (*model.Customer, error) {
	return nil, errors.New("not implemented")
}

func (m *MockCustomerRepoError) FindByDocumentNumber(ctx context.Context, id string) (*model.Customer, error) {
	return nil, errors.New("erro ao buscar customer")
}
