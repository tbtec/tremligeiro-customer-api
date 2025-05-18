package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type CreateCustomerController struct {
	container *container.Container
	usc       *usecase.UscCreateCustomer
}

func NewCreateCustomerController(container *container.Container) *CreateCustomerController {
	return &CreateCustomerController{
		container: container,
		usc: usecase.NewUseCaseCreateCustomer(
			gateway.NewCustomerGateway(container.CustomerRepository),
			presenter.NewCustomerPresenter(),
		),
	}
}

func (ctl *CreateCustomerController) Execute(ctx context.Context, customer dto.CreateCustomer) (dto.Customer, error) {
	return ctl.usc.Create(ctx, customer)
}
