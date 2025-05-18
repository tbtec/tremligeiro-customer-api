// go
package dto

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestCreateCustomerStruct(t *testing.T) {
	c := CreateCustomer{
		Name:           "John Doe",
		DocumentNumber: "123456789",
		Email:          "john@example.com",
	}
	if c.Name != "John Doe" {
		t.Errorf("expected Name to be 'John Doe', got %s", c.Name)
	}
	if c.DocumentNumber != "123456789" {
		t.Errorf("expected DocumentNumber to be '123456789', got %s", c.DocumentNumber)
	}
	if c.Email != "john@example.com" {
		t.Errorf("expected Email to be 'john@example.com', got %s", c.Email)
	}
}

func TestUpdateCustomerStruct(t *testing.T) {
	now := time.Now()
	u := UpdateCustomer{
		CustomerId:     "id1",
		Name:           "Jane Doe",
		DocumentNumber: "987654321",
		Email:          "jane@example.com",
		UpdatedAt:      now,
	}
	if u.CustomerId != "id1" {
		t.Errorf("expected CustomerId to be 'id1', got %s", u.CustomerId)
	}
	if u.Name != "Jane Doe" {
		t.Errorf("expected Name to be 'Jane Doe', got %s", u.Name)
	}
	if u.DocumentNumber != "987654321" {
		t.Errorf("expected DocumentNumber to be '987654321', got %s", u.DocumentNumber)
	}
	if u.Email != "jane@example.com" {
		t.Errorf("expected Email to be 'jane@example.com', got %s", u.Email)
	}
	if !u.UpdatedAt.Equal(now) {
		t.Errorf("expected UpdatedAt to be %v, got %v", now, u.UpdatedAt)
	}
}

func TestFindCustomerStruct(t *testing.T) {
	f := FindCustomer{
		DocumentNumber: "123456789",
	}
	if f.DocumentNumber != "123456789" {
		t.Errorf("expected DocumentNumber to be '123456789', got %s", f.DocumentNumber)
	}
}

func TestFindOneCustomerStruct(t *testing.T) {
	f := FindOneCustomer{
		CustomerId: "id2",
	}
	if f.CustomerId != "id2" {
		t.Errorf("expected CustomerId to be 'id2', got %s", f.CustomerId)
	}
}

func TestCustomerStruct(t *testing.T) {
	now := time.Now()
	c := Customer{
		CustomerId:     "id3",
		Name:           "Alice",
		DocumentNumber: "111222333",
		Email:          "alice@example.com",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	if c.CustomerId != "id3" {
		t.Errorf("expected CustomerId to be 'id3', got %s", c.CustomerId)
	}
	if c.Name != "Alice" {
		t.Errorf("expected Name to be 'Alice', got %s", c.Name)
	}
	if c.DocumentNumber != "111222333" {
		t.Errorf("expected DocumentNumber to be '111222333', got %s", c.DocumentNumber)
	}
	if c.Email != "alice@example.com" {
		t.Errorf("expected Email to be 'alice@example.com', got %s", c.Email)
	}
	if !c.CreatedAt.Equal(now) {
		t.Errorf("expected CreatedAt to be %v, got %v", now, c.CreatedAt)
	}
	if !c.UpdatedAt.Equal(now) {
		t.Errorf("expected UpdatedAt to be %v, got %v", now, c.UpdatedAt)
	}
}

func TestCustomerContentStruct(t *testing.T) {
	c := Customer{
		CustomerId:     "id4",
		Name:           "Bob",
		DocumentNumber: "444555666",
		Email:          "bob@example.com",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	cc := CustomerContent{
		Content: c,
	}
	if !reflect.DeepEqual(cc.Content, c) {
		t.Errorf("expected Content to be %+v, got %+v", c, cc.Content)
	}
}

func TestCustomerJSONTags(t *testing.T) {
	c := Customer{
		CustomerId:     "id5",
		Name:           "Carol",
		DocumentNumber: "777888999",
		Email:          "carol@example.com",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	}
	data, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("failed to marshal Customer: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("failed to unmarshal Customer JSON: %v", err)
	}
	expectedKeys := []string{"id", "name", "documentNumber", "email", "createdAt", "updatedAt"}
	for _, key := range expectedKeys {
		if _, ok := m[key]; !ok {
			t.Errorf("expected JSON to have key %s", key)
		}
	}
}