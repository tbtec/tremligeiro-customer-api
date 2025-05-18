package gateway

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

// MockCustomerRepository mock da interface ICustomerRepository
type MockCustomerRepository struct {
	mock.Mock
}

// err := gtw.customerRepository.Create(ctx, &customerModel)
func (m *MockCustomerRepository) Create(ctx context.Context, customer *model.Customer) error {
	args := m.Called(ctx, customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Customer, error) {
	args := m.Called(ctx, documentNumber)
	if c := args.Get(0); c != nil {
		return c.(*model.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCustomerRepository) FindOne(ctx context.Context, id string) (*model.Customer, error) {
	args := m.Called(ctx, id)
	if c := args.Get(0); c != nil {
		return c.(*model.Customer), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCustomerGateway_Create_Success(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	gateway := NewCustomerGateway(mockRepo)

	ctx := context.Background()
	customer := &entity.Customer{
		ID:             "123",
		Name:           "João",
		Email:          "joao@example.com",
		DocumentNumber: "123456789",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	mockRepo.On("Create", ctx, mock.MatchedBy(func(c *model.Customer) bool {
		return c.ID == customer.ID &&
			c.Name == customer.Name &&
			c.Email == customer.Email &&
			c.DocumentNumber == customer.DocumentNumber
	})).Return(nil)

	err := gateway.Create(ctx, customer)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerGateway_Create_Error(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	gateway := NewCustomerGateway(mockRepo)

	ctx := context.Background()
	customer := &entity.Customer{ID: "123"}

	mockRepo.On("Create", ctx, mock.Anything).Return(errors.New("db error"))

	err := gateway.Create(ctx, customer)

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestCustomerGateway_FindByDocumentNumber_Success(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	gateway := NewCustomerGateway(mockRepo)

	ctx := context.Background()
	documentNumber := "123456789"

	customerModel := &model.Customer{
		ID:             "123",
		Name:           "João",
		Email:          "joao@example.com",
		DocumentNumber: documentNumber,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	mockRepo.On("FindByDocumentNumber", ctx, documentNumber).Return(customerModel, nil)

	customer, err := gateway.FindByDocumentNumber(ctx, documentNumber)

	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, customerModel.ID, customer.ID)
	assert.Equal(t, customerModel.Name, customer.Name)
	mockRepo.AssertExpectations(t)
}

func TestCustomerGateway_FindByDocumentNumber_NotFound(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	gateway := NewCustomerGateway(mockRepo)

	ctx := context.Background()
	documentNumber := "000000000"

	mockRepo.On("FindByDocumentNumber", ctx, documentNumber).Return(nil, errors.New("not found"))

	customer, err := gateway.FindByDocumentNumber(ctx, documentNumber)

	assert.Error(t, err)
	assert.Nil(t, customer)
	mockRepo.AssertExpectations(t)
}

func TestCustomerGateway_FindOne_Success(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	gateway := NewCustomerGateway(mockRepo)

	ctx := context.Background()
	id := "123"

	customerModel := &model.Customer{
		ID:             id,
		Name:           "Maria",
		Email:          "maria@example.com",
		DocumentNumber: "987654321",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	mockRepo.On("FindOne", ctx, id).Return(customerModel, nil)

	customer, err := gateway.FindOne(ctx, id)

	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, customerModel.ID, customer.ID)
	assert.Equal(t, customerModel.Name, customer.Name)
	mockRepo.AssertExpectations(t)
}

func TestCustomerGateway_FindOne_NotFound(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	gateway := NewCustomerGateway(mockRepo)

	ctx := context.Background()
	id := "999"

	mockRepo.On("FindOne", ctx, id).Return(nil, nil)

	customer, err := gateway.FindOne(ctx, id)

	assert.NoError(t, err)
	assert.Nil(t, customer)
	mockRepo.AssertExpectations(t)
}

func TestCustomerGateway_FindOne_Error(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	gateway := NewCustomerGateway(mockRepo)

	ctx := context.Background()
	id := "999"

	mockRepo.On("FindOne", ctx, id).Return(nil, errors.New("db error"))

	customer, err := gateway.FindOne(ctx, id)

	assert.Error(t, err)
	assert.Nil(t, customer)
	mockRepo.AssertExpectations(t)
}
