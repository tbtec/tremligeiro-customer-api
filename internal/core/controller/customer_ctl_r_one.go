package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type FindOneCustomerController struct {
	usc *usecase.UscFindCustomer
}

func NewFindOneCustomerController(container *container.Container) *FindOneCustomerController {
	return &FindOneCustomerController{
		usc: usecase.NewUseCaseFindOneCustomer(
			gateway.NewCustomerGateway(container.CustomerRepository),
		),
	}
}

func (ctl *FindOneCustomerController) Execute(ctx context.Context, input dto.FindOneCustomer) (dto.CustomerContent, error) {
	return ctl.usc.FindOne(ctx, input)
}
