package application

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

//go:generate mockery --name ProductInterface
//go:generate mockery --name ProductServiceInterface
//go:generate mockery --name ProductPersistenceInterface

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `validate:"required,uuid"`
	Name   string  `validate:"required"`
	Price  float64 `validate:"required,gte=0"`
	Status string  `validate:"required,oneof='disabled' 'enabled' "`
}

func NewProduct(name *string, price *float64) (product *Product) {
	newUUID, _ := uuid.NewUUID()
	product = &Product{
		ID:     newUUID.String(),
		Status: DISABLED,
	}
	if name != nil {
		product.Name = *name
	}
	if price != nil {
		product.Price = *price
	}
	if valid, _ := product.IsValid(); valid {
		_ = product.Enable()
	}
	return product
}

func (p *Product) IsValid() (bool, error) {
	err := validator.New().Struct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("price must be greater than zero to enable the product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("price must be equal zero to disable the product")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
