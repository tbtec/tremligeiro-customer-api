package controller

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
	"github.com/tbtec/tremligeiro/test/repository"
)

func TestHandle_Success(t *testing.T) {

	mockCustomer := model.Customer{
		ID:             "customerID",
		Name:           "Test Customer",
		DocumentNumber: "123456789",
		Email:          "teste@teste.com",
	}

	container := &container.Container{
		CustomerRepository: &repository.MockCustomerRepo{
			FindByDocumentNumberFunc: func(ctx context.Context, id string) (*model.Customer, error) {
				return &mockCustomer, nil
			},
		},
	}
	ctrl := NewCustomerFindRestController(container)

	req := httpserver.Request{
		Query: map[string]string{"documentNumber": "123456789"},
	}
	resp := ctrl.Handle(context.Background(), req)

	assert.Equal(t, 200, resp.Code)
}

func TestHandle_Invalid(t *testing.T) {

	container := &container.Container{
		CustomerRepository: &repository.MockCustomerRepo{
			FindByDocumentNumberFunc: func(ctx context.Context, id string) (*model.Customer, error) {
				return nil, errors.New("record not found")
			},
		},
	}
	ctrl := NewCustomerFindRestController(container)

	req := httpserver.Request{
		Query: map[string]string{"documentNumber": "123456789"},
	}
	resp := ctrl.Handle(context.Background(), req)

	assert.Equal(t, httpserver.HandleError(context.Background(), &strconv.NumError{
		Func: "Atoi",
		Num:  "123456789",
		Err:  errors.New("invalid syntax"),
	}).Code, resp.Code)
}
