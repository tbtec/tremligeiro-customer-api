package controller

import (
	"context"
	"fmt"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type CustomerFindOneController struct {
	controller *ctl.FindOneCustomerController
}

type CustomerFindOneResponse struct {
	CustomerId     string `json:"id"`
	Name           string `json:"name"`
	DocumentNumber string `json:"documentNumber"`
	Email          string `json:"email"`
}

func NewCustomerFindOneRestController(container *container.Container) httpserver.IController {
	return &CustomerFindOneController{
		controller: ctl.NewFindOneCustomerController(container),
	}
}

func (ctl *CustomerFindOneController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	customerIdRequest := dto.FindOneCustomer{
		CustomerId: request.Params["customerId"],
	}

	fmt.Println("command: ", customerIdRequest)

	output, err := ctl.controller.Execute(ctx, customerIdRequest)

	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(output)
}

func buildCustomerFindOneResponse(customer *entity.Customer) CustomerFindOneResponse {
	return CustomerFindOneResponse{
		CustomerId:     customer.ID,
		Name:           customer.Name,
		DocumentNumber: customer.DocumentNumber,
		Email:          customer.Email,
	}
}
