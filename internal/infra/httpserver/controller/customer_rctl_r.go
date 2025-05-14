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

type CustomerFindController struct {
	controller *ctl.FindCustomerController
}

type CustomerFindResponse struct {
	CustomerId     string `json:"id"`
	Name           string `json:"name"`
	DocumentNumber string `json:"documentNumber"`
	Email          string `json:"email"`
}

func NewCustomerFindRestController(container *container.Container) httpserver.IController {
	return &CustomerFindController{
		controller: ctl.NewFindCustomerController(container),
	}
}

func (ctl *CustomerFindController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	documentNumberRequest := dto.FindCustomer{
		DocumentNumber: request.Query["documentNumber"],
	}

	fmt.Println("command: ", documentNumberRequest)

	output, err := ctl.controller.Execute(ctx, documentNumberRequest)

	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(output)
}

func buildCustomerFindResponse(customer *entity.Customer) CustomerFindResponse {
	return CustomerFindResponse{
		CustomerId:     customer.ID,
		Name:           customer.Name,
		DocumentNumber: customer.DocumentNumber,
		Email:          customer.Email,
	}
}
