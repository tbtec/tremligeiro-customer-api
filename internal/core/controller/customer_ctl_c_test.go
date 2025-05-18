package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/test/repository"
)

func TestCreateCustomerController_Execute_Success(t *testing.T) {
	ctx := context.Background()

	customerRepo := &repository.MockCustomerRepo{
		CreateFunc: func(ctx context.Context, customer *model.Customer) error {
			return nil
		},
	}

	testContainer := &container.Container{
		CustomerRepository: customerRepo,
	}

	controller := NewCreateCustomerController(testContainer)

	input := dto.CreateCustomer{
		Name:           "Customer 1",
		DocumentNumber: "123456789",
		Email:          "teste@teste.com",
	}

	result, err := controller.Execute(ctx, input)
	assert.NoError(t, err)
	assert.Equal(t, "Customer 1", result.Name)

}
