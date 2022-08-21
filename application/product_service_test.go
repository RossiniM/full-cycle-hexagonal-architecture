package application_test

import (
	"github.com/RossiniM/full-cycle-hexagonal-architecture/application"
	"github.com/RossiniM/full-cycle-hexagonal-architecture/application/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {

	product := &mocks.ProductInterface{}
	persistence := &mocks.ProductPersistenceInterface{}
	service := application.ProductService{Persistence: persistence}
	persistence.On("Get", mock.Anything).Return(product, nil)

	result, err := service.Get("abc")

	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	name := "product-1"
	price := 50.0
	product := application.NewProduct(&name, &price)

	persistence := &mocks.ProductPersistenceInterface{}

	matcher := mock.MatchedBy(func(p application.ProductInterface) bool {
		require.Equal(t, product.Name, p.GetName())
		require.Equal(t, product.Price, p.GetPrice())
		require.Equal(t, product.Status, p.GetStatus())
		return true
	})

	persistence.On("Save", matcher).Return(product, nil)
	service := application.ProductService{Persistence: persistence}

	result, err := service.Create("product-1", 50)

	require.Nil(t, err)
	require.NotNil(t, result)

}

func TestProductService_Disable(t *testing.T) {
	name := "product-1"
	price := 50.0
	product := application.NewProduct(&name, &price)
	require.Equal(t, application.ENABLED, product.Status)
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)
}

func TestProductService_Enable(t *testing.T) {

	product := application.NewProduct(nil, nil)
	require.Equal(t, application.DISABLED, product.Status)
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.Status)
}
