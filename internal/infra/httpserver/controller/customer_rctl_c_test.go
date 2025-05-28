package controller

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
	"github.com/tbtec/tremligeiro/test/repository"
)

func newMockContainer() *container.Container {
	return &container.Container{
		CustomerRepository: &repository.MockCustomerRepoInterface{},
	}
}

func TestNewCustomerCreateRestController(t *testing.T) {
	container := newMockContainer()
	ctrl := NewCustomerCreateRestController(container)
	assert.NotNil(t, ctrl)
}

func TestCustomerCreateRestController_Handle(t *testing.T) {
	container := newMockContainer()
	ctrl := NewCustomerCreateRestController(container)

	input := dto.CreateCustomer{
		Name:           "Customer 1",
		DocumentNumber: "123456789",
		Email:          "teste@teste.com",
	}

	inputBytes, _ := json.Marshal(input)
	req := httpserver.Request{Body: inputBytes}

	resp := ctrl.Handle(context.Background(), req)

	assert.Equal(t, 200, resp.Code)

	var output struct {
		Status string `json:"status"`
	}

	bodyBytes, err := json.Marshal(resp.Body)
	assert.NoError(t, err)
	err = json.Unmarshal(bodyBytes, &output)
	assert.NoError(t, err)
	assert.Equal(t, "", output.Status)
}

func TestCustomerCreateRestController_Handle_InvalidJSON(t *testing.T) {
	container := newMockContainer()
	ctrl := NewCustomerCreateRestController(container)

	req := httpserver.Request{Body: []byte(`{invalid json}`)}

	resp := ctrl.Handle(context.Background(), req)

	assert.NotEqual(t, 200, resp.Code)
}

func TestCustomerCreateRestController_Handle_MissingFields(t *testing.T) {
	container := newMockContainer()
	ctrl := NewCustomerCreateRestController(container)

	input := map[string]interface{}{
		"Name": "",
	}
	inputBytes, _ := json.Marshal(input)
	req := httpserver.Request{Body: inputBytes}

	resp := ctrl.Handle(context.Background(), req)

	assert.NotEqual(t, 200, resp.Code)
}

func TestCustomerCreateRestController_Handle_ExecuteError(t *testing.T) {
	container := &container.Container{
		CustomerRepository: &repository.MockCustomerRepoError{},
	}
	ctrl := NewCustomerCreateRestController(container)

	input := dto.CreateCustomer{
		Name:           "Customer 1",
		DocumentNumber: "123456789",
		Email:          "teste@teste.com",
	}
	inputBytes, _ := json.Marshal(input)
	req := httpserver.Request{Body: inputBytes}

	resp := ctrl.Handle(context.Background(), req)

	assert.NotEqual(t, 200, resp.Code)

	errMsg, ok := resp.Body.(httpserver.ErrorMessage)
	assert.True(t, ok)
	assert.Contains(t, errMsg.Error.Description, "Internal Server Error")
}
