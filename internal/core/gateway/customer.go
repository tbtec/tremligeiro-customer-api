package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
)

type CustomerGateway struct {
	customerRepository repository.ICustomerRepository
}

func NewCustomerGateway(customerRepository repository.ICustomerRepository) *CustomerGateway {
	return &CustomerGateway{
		customerRepository: customerRepository,
	}
}

func (gtw *CustomerGateway) Create(ctx context.Context, customer *entity.Customer) error {

	customerModel := model.Customer{
		ID:             customer.ID,
		Name:           customer.Name,
		Email:          customer.Email,
		DocumentNumber: customer.DocumentNumber,
		CreatedAt:      customer.CreatedAt,
		UpdatedAt:      customer.UpdatedAt,
	}

	err := gtw.customerRepository.Create(ctx, &customerModel)

	if err != nil {
		return err
	}

	return nil
}

func (gtw *CustomerGateway) FindByDocumentNumber(ctx context.Context, documentNumber string) (*entity.Customer, error) {
	customerModel, err := gtw.customerRepository.FindByDocumentNumber(ctx, documentNumber)
	if err != nil {
		return nil, err
	}

	customer := entity.Customer{
		ID:             customerModel.ID,
		Name:           customerModel.Name,
		DocumentNumber: customerModel.DocumentNumber,
		Email:          customerModel.Email,
		CreatedAt:      customerModel.CreatedAt,
		UpdatedAt:      customerModel.UpdatedAt,
	}

	return &customer, nil
}

func (gtw *CustomerGateway) FindOne(ctx context.Context, id string) (*entity.Customer, error) {
	customer := entity.Customer{}

	customerModel, err := gtw.customerRepository.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	if customerModel != nil {
		customer = entity.Customer{
			ID:             customerModel.ID,
			Name:           customerModel.Name,
			DocumentNumber: customerModel.DocumentNumber,
			Email:          customerModel.Email,
			CreatedAt:      customerModel.CreatedAt,
			UpdatedAt:      customerModel.UpdatedAt,
		}
		return &customer, nil
	}

	return nil, nil
}
