package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type UscFindOneCustomer struct {
	customerGateway   *gateway.CustomerGateway
	customerPresenter *presenter.CustomerPresenter
}

func NewUseCaseFindOneCustomer(customerGateway *gateway.CustomerGateway) *UscFindCustomer {
	return &UscFindCustomer{
		customerGateway:   customerGateway,
		customerPresenter: presenter.NewCustomerPresenter(),
	}
}

func (usc *UscFindCustomer) FindOne(ctx context.Context, findCustomer dto.FindOneCustomer) (dto.CustomerContent, error) {
	customer, err := usc.customerGateway.FindOne(ctx, findCustomer.CustomerId)
	if err != nil {
		return dto.CustomerContent{}, err
	}

	return usc.customerPresenter.BuildCustomerContentResponse(*customer), nil
}
