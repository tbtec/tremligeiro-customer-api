package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type UscFindCustomer struct {
	customerGateway   *gateway.CustomerGateway
	customerPresenter *presenter.CustomerPresenter
}

func NewUseCaseFindCustomer(customerGateway *gateway.CustomerGateway) *UscFindCustomer {
	return &UscFindCustomer{
		customerGateway:   customerGateway,
		customerPresenter: presenter.NewCustomerPresenter(),
	}
}

func (usc *UscFindCustomer) FindByDocumentNumber(ctx context.Context, findCustomer dto.FindCustomer) (dto.CustomerContent, error) {
	customer, err := usc.customerGateway.FindByDocumentNumber(ctx, findCustomer.DocumentNumber)
	if err != nil {
		return dto.CustomerContent{}, err
	}

	return usc.customerPresenter.BuildCustomerContentResponse(*customer), nil
}
