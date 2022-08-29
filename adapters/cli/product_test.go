package cli

import (
	"errors"
	"fmt"
	"github.com/RossiniM/full-cycle-hexagonal-architecture/application"
	"github.com/RossiniM/full-cycle-hexagonal-architecture/application/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun_CreateOK(t *testing.T) {

	productServiceInterface := &mocks.ProductServiceInterface{}
	name := "product-1"
	price := 5.9
	product := application.NewProduct(&name, &price)
	logExpected := "Product ID %s with name %s has been created with the price %f and status enabled"

	productServiceInterface.
		On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("float64")).
		Return(product, nil)
	result, err := Run(productServiceInterface, "create", "1L", "product-1", 5.0)

	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf(logExpected, product.ID, product.Name, product.Price), result)
}

func TestRun_CreateFailed(t *testing.T) {

	productServiceInterface := &mocks.ProductServiceInterface{}

	errorExpected := errors.New("unexpected error at product creation")
	productServiceInterface.
		On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("float64")).
		Return(nil, errorExpected)
	_, err := Run(productServiceInterface, "create", "1L", "product-1", 5.0)

	require.NotNil(t, err)
	require.Equal(t, errorExpected.Error(), err.Error())
}

func TestRun_Enabled_OK(t *testing.T) {
	productServiceInterface := &mocks.ProductServiceInterface{}
	name := "product-1"
	product := application.NewProduct(&name, nil)
	product.Price = 5.0
	logExpected := "Product ID %s with name %s  with the price %f has been enabled"
	productServiceInterface.
		On("Get", mock.AnythingOfType("string")).
		Return(product, nil)

	result, err := Run(productServiceInterface, "enable", product.ID, product.Name, product.Price)

	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf(logExpected, product.ID, product.Name, product.Price), result)
}

func TestRun_EnableFailed(t *testing.T) {

	productServiceInterface := &mocks.ProductServiceInterface{}
	name := "product-1"
	product := application.NewProduct(&name, nil)
	errorExpected := "price must be greater than zero to enable the product"
	productServiceInterface.
		On("Get", mock.AnythingOfType("string")).
		Return(product, nil)

	_, err := Run(productServiceInterface, "enable", product.ID, product.Name, product.Price)

	require.NotNil(t, err)
	require.Equal(t, errorExpected, err.Error())
}

func TestRun_Disabled_OK(t *testing.T) {
	productServiceInterface := &mocks.ProductServiceInterface{}
	name := "product-1"
	product := application.NewProduct(&name, nil)
	product.Price = 5.0
	product.Enable()
	product.Price = 0.0
	logExpected := "Product ID %s with name %s  with the price %f has been disabled"
	productServiceInterface.
		On("Get", mock.AnythingOfType("string")).
		Return(product, nil)

	result, err := Run(productServiceInterface, "disable", product.ID, product.Name, product.Price)

	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf(logExpected, product.ID, product.Name, product.Price), result)
}

func TestRun_DisableFailed(t *testing.T) {

	productServiceInterface := &mocks.ProductServiceInterface{}
	name := "product-1"
	product := application.NewProduct(&name, nil)
	product.Price = 5.0
	product.Enable()
	errorExpected := "price must be equal zero to disable the product"
	productServiceInterface.
		On("Get", mock.AnythingOfType("string")).
		Return(product, nil)

	_, err := Run(productServiceInterface, "disable", product.ID, product.Name, product.Price)

	require.NotNil(t, err)
	require.Equal(t, errorExpected, err.Error())
}

func TestRun_Default(t *testing.T) {

	productServiceInterface := &mocks.ProductServiceInterface{}
	name := "product-1"
	product := application.NewProduct(&name, nil)

	logExpected := fmt.Sprintf(
		"Product ID %s with name %s with the price %f and status %s",
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus())

	productServiceInterface.
		On("Get", mock.AnythingOfType("string")).
		Return(product, nil)

	result, err := Run(productServiceInterface, "invalid", product.ID, product.Name, product.Price)

	require.Nil(t, err)
	require.Equal(t, logExpected, result)
}
