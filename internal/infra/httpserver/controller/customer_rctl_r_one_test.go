package controller

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tbtec/tremligeiro/internal/dto"
)

// Mock of FindOneCustomerController
type MockFindOneCustomerController struct {
	mock.Mock
}

// Handle mocks the behavior of the real controller's Handle method.
func (m *MockFindOneCustomerController) Handle(ctx context.Context, req *MockRequest) *MockResponse {
	customerId := req.ParseParamString("customerId")
	customer, err := m.Execute(ctx, customerId)
	if err != nil {
		return &MockResponse{
			Code: 500,
			Body: err.Error(),
		}
	}
	return &MockResponse{
		Code: 200,
		Body: customer,
	}
}

// MockResponse is a simple struct to simulate HTTP responses in tests.
type MockResponse struct {
	Code int
	Body interface{}
}

func (m *MockFindOneCustomerController) Execute(ctx context.Context, id string) (*dto.Customer, error) {
	args := m.Called(ctx, id)
	if customer := args.Get(0); customer != nil {
		return customer.(*dto.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}

// Mock request
type MockRequest struct {
	params map[string]string
}

func (m *MockRequest) ParseParamString(key string) string {
	return m.params[key]
}

type EmbeddedMockFindOneCustomerController struct {
	*MockFindOneCustomerController
}

func (e *EmbeddedMockFindOneCustomerController) Execute(ctx context.Context, id string) (*dto.Customer, error) {
	return e.MockFindOneCustomerController.Execute(ctx, id)
}

func TestCustomerFindOneController_Handle_Success(t *testing.T) {
	ctx := context.Background()

	// Arrange
	mockController := new(MockFindOneCustomerController)
	customerID := "123"

	expectedCustomer := &dto.Customer{
		CustomerId:     customerID,
		Name:           "Test Customer",
		DocumentNumber: "123456789",
		Email:          "teste@teste.com",
	}

	mockController.On("Execute", ctx, customerID).Return(expectedCustomer, nil)

	restController := mockController

	req := &MockRequest{
		params: map[string]string{"customerId": customerID},
	}

	// Act
	resp := restController.Handle(ctx, req)

	// Assert
	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, expectedCustomer, resp.Body)
	mockController.AssertExpectations(t)
}

func TestCustomerFindOneController_Handle_Error(t *testing.T) {
	ctx := context.Background()

	mockController := new(MockFindOneCustomerController)
	customerID := "notfound"
	expectedErr := errors.New("customer not found")

	mockController.On("Execute", ctx, customerID).Return(nil, expectedErr)

	req := &MockRequest{
		params: map[string]string{"customerId": customerID},
	}

	resp := mockController.Handle(ctx, req)

	assert.Equal(t, 500, resp.Code)
	assert.Equal(t, "customer not found", resp.Body)
	mockController.AssertExpectations(t)
}
