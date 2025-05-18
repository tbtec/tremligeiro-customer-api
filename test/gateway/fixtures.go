package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type MockCustomerGateway struct {
	CreateFunc func(ctx context.Context, customer *entity.Customer) error
}

type MockCustomerPresenter struct {
	BuildProductCreateResponseFunc func(customer entity.Customer) dto.Customer
}

// func (gtw *CustomerGateway) Create(ctx context.Context, customer *entity.Customer) error {
func (m *MockCustomerGateway) Create(ctx context.Context, customer *entity.Customer) error {

	return m.CreateFunc(ctx, customer)
}

func (m *MockCustomerPresenter) BuildCustomerCreateResponse(customer entity.Customer) dto.Customer {
	return m.BuildProductCreateResponseFunc(customer)
}
