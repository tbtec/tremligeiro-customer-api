package presenter_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
)

func TestBuildCustomerCreateResponse(t *testing.T) {
	// Arrange: Criar um customer de exemplo
	customer := entity.Customer{
		ID:             "1",
		Name:           "John Doe",
		DocumentNumber: "123456789",
		Email:          "john.doe@example.com",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	// Instanciar o presenter
	presenter := presenter.NewCustomerPresenter()

	// Act: Chamar o método BuildCustomerCreateResponse
	result := presenter.BuildCustomerCreateResponse(customer)

	// Assert: Verificar se o resultado está correto
	assert.Equal(t, customer.ID, result.CustomerId)
	assert.Equal(t, customer.Name, result.Name)
	assert.Equal(t, customer.DocumentNumber, result.DocumentNumber)
	assert.Equal(t, customer.Email, result.Email)
	assert.Equal(t, customer.CreatedAt, result.CreatedAt)
	assert.Equal(t, customer.UpdatedAt, result.UpdatedAt)
}

func TestBuildCustomerContentResponse(t *testing.T) {
	// Arrange: Criar um customer de exemplo
	customer := entity.Customer{
		ID:             "1",
		Name:           "John Doe",
		DocumentNumber: "123456789",
		Email:          "john.doe@example.com",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	// Instanciar o presenter
	presenter := presenter.NewCustomerPresenter()

	// Act: Chamar o método BuildCustomerContentResponse
	result := presenter.BuildCustomerContentResponse(customer)

	// Assert: Verificar se o resultado tem o campo Content correto
	assert.Equal(t, customer.ID, result.Content.CustomerId)
	assert.Equal(t, customer.Name, result.Content.Name)
	assert.Equal(t, customer.DocumentNumber, result.Content.DocumentNumber)
	assert.Equal(t, customer.Email, result.Content.Email)
	assert.Equal(t, customer.CreatedAt, result.Content.CreatedAt)
	assert.Equal(t, customer.UpdatedAt, result.Content.UpdatedAt)
}
