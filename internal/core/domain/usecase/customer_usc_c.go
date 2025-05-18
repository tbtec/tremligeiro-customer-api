package usecase

import (
	"context"
	"time"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/types/ulid"
)

type UscCreateCustomer struct {
	customerGateway   *gateway.CustomerGateway
	customerPresenter *presenter.CustomerPresenter
}

func NewUseCaseCreateCustomer(customerGateway *gateway.CustomerGateway,
	customerPresenter *presenter.CustomerPresenter) *UscCreateCustomer {
	return &UscCreateCustomer{
		customerGateway:   customerGateway,
		customerPresenter: customerPresenter,
	}
}

func (usc *UscCreateCustomer) Create(ctx context.Context, customerDto dto.CreateCustomer) (dto.Customer, error) {

	customer := entity.Customer{
		ID:             ulid.NewUlid().String(),
		Name:           customerDto.Name,
		DocumentNumber: customerDto.DocumentNumber,
		Email:          customerDto.Email,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}

	err := usc.customerGateway.Create(ctx, &customer)
	if err != nil {
		return dto.Customer{}, err
	}

	return usc.customerPresenter.BuildCustomerCreateResponse(customer), nil
}
