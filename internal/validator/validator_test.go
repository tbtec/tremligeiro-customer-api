package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/types/xerrors"
)

type testStruct struct {
	Name string `validate:"required"`
	Age  int    `validate:"min=1"`
}

func TestNewValidator(t *testing.T) {
	v := newValidator()
	assert.NotNil(t, v)
}

func TestValidate_Success(t *testing.T) {
	ts := testStruct{Name: "John", Age: 10}
	err := Validate(ts)
	assert.NoError(t, err)
}

func TestValidate_RequiredError(t *testing.T) {
	ts := testStruct{Age: 10}
	err := Validate(ts)
	assert.Error(t, err)
	verr, ok := err.(xerrors.ValidationError)
	assert.True(t, ok)
	assert.True(t, len(verr.Fields) > 0)
}

func TestValidate_MinError(t *testing.T) {
	ts := testStruct{Name: "John", Age: 0}
	err := Validate(ts)
	assert.Error(t, err)
	verr, ok := err.(xerrors.ValidationError)
	assert.True(t, ok)
	assert.True(t, len(verr.Fields) > 0)
}

/*
func TestAdapt_Required(t *testing.T) {
	v := validator.New()
	type S struct {
		Foo string `validate:"required"`
	}
	s := S{}
	err := v.Struct(s)
	verr := adapt(err)
	assert.Contains(t, verr.Error(), "Invalid Body")
}

func TestAdapt_OtherTag(t *testing.T) {
	v := validator.New()
	type S struct {
		Foo int `validate:"min=10"`
	}
	s := S{Foo: 1}
	err := v.Struct(s)
	verr := adapt(err)
	assert.Contains(t, verr.Error(), "Invalid Body")
}*/

func TestExtract(t *testing.T) {
	assert.Equal(t, "Foo", extract("S.Foo"))
	assert.Equal(t, "S", extract("S"))
	assert.Equal(t, "Bar.Baz", extract("A.Bar.Baz"))
}
