package controller

import (
	"context"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
	"github.com/tbtec/tremligeiro/internal/validator"
)

type CustomerCreateRestController struct {
	controller *ctl.CreateCustomerController
}

type CustomerCreateRequest struct {
	Name           string `json:"name" validate:"required"`
	DocumentNumber string `json:"documentNumber" validate:"required"`
	Email          string `json:"email" validate:"required"`
}

type CustomerCreateResponse struct {
	CustomerId     string `json:"id"`
	Name           string `json:"name"`
	DocumentNumber string `json:"documentNumber"`
	Email          string `json:"email"`
}

func NewCustomerCreateRestController(container *container.Container) httpserver.IController {
	return &CustomerCreateRestController{
		controller: ctl.NewCreateCustomerController(container),
	}
}

func (ctl *CustomerCreateRestController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	pustomerRequest := dto.CreateCustomer{}

	errBody := request.ParseBody(ctx, &pustomerRequest)
	if errBody != nil {
		return httpserver.HandleError(ctx, errBody)
	}

	err := validator.Validate(pustomerRequest)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	output, err := ctl.controller.Execute(ctx, pustomerRequest)

	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(output)
}
