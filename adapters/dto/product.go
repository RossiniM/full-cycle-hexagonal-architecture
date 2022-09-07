package dto

import "github.com/RossiniM/full-cycle-hexagonal-architecture/application"

type ProductRequest struct {
	ID     string  `json:"id" `
	Name   string  `json:"name" `
	Price  float64 `json:"price" `
	Status string  `json:"status"`
}

func NewProductRequest() *ProductRequest {
	return &ProductRequest{}
}

func (p *ProductRequest) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}
	return product, nil
}
