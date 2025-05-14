package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type FindCustomerController struct {
	usc *usecase.UscFindCustomer
}

func NewFindCustomerController(container *container.Container) *FindCustomerController {
	return &FindCustomerController{
		usc: usecase.NewUseCaseFindCustomer(
			gateway.NewCustomerGateway(container.CustomerRepository),
		),
	}
}

func (ctl *FindCustomerController) Execute(ctx context.Context, input dto.FindCustomer) (dto.CustomerContent, error) {
	return ctl.usc.FindByDocumentNumber(ctx, input)
}
