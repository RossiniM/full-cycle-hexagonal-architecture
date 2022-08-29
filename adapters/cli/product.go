package cli

import (
	"fmt"
	"github.com/RossiniM/full-cycle-hexagonal-architecture/application"
)

func Run(service application.ProductServiceInterface, action string,
	productID string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf(
			"Product ID %s with name %s has been created with the price %f and status %s",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())
	case "enable":
		product, err := service.Get(productID)
		if err != nil {
			return "", err
		}
		err = product.Enable()
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf(
			"Product ID %s with name %s  with the price %f has been enabled",
			product.GetID(),
			product.GetName(),
			product.GetPrice())

	case "disable":
		product, err := service.Get(productID)
		if err != nil {
			return "", err
		}
		err = product.Disable()
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf(
			"Product ID %s with name %s  with the price %f has been disabled",
			product.GetID(),
			product.GetName(),
			product.GetPrice())
	default:
		product, err := service.Get(productID)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf(
			"Product ID %s with name %s with the price %f and status %s",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())
	}
	fmt.Printf(result)
	return result, nil
}
