package application

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable_OK(t *testing.T) {
	product := Product{
		ID:     "1",
		Name:   "P1",
		Price:  5,
		Status: "",
	}
	err := product.Enable()
	require.Nil(t, err)
	require.Equal(t, ENABLED, product.Status)
}

func TestProduct_Enable_Error(t *testing.T) {
	product := Product{
		ID:     "1",
		Name:   "P1",
		Price:  0,
		Status: "",
	}
	err := product.Enable()
	require.Error(t, err, "price must be greater than zero to enable the product")
	require.Equal(t, "", product.Status)
}

func TestProduct_Disable_OK(t *testing.T) {
	product := Product{
		ID:     "1",
		Name:   "P1",
		Price:  0,
		Status: "",
	}
	err := product.Disable()
	require.Nil(t, err)
	require.Equal(t, DISABLED, product.Status)
}

func TestProduct_Disable_Error(t *testing.T) {
	product := Product{
		ID:     "1",
		Name:   "P1",
		Price:  5,
		Status: "",
	}
	err := product.Disable()
	require.Error(t, err, "price must be equal zero to disable the product")
	require.Equal(t, "", product.Status)
}

func TestProduct_isValid_True(t *testing.T) {
	newUUID, _ := uuid.NewUUID()
	product := Product{
		ID:     newUUID.String(),
		Name:   "P1",
		Price:  5,
		Status: DISABLED,
	}
	isValid, err := product.isValid()
	require.Nil(t, err)
	require.True(t, isValid)
}

func TestProduct_isValid_False(t *testing.T) {
	newUUID, _ := uuid.NewUUID()
	product := Product{
		ID:     newUUID.String(),
		Name:   "",
		Price:  -4,
		Status: "DLED",
	}
	isValid, err := product.isValid()

	require.NotNil(t, err)
	require.False(t, isValid)
}
