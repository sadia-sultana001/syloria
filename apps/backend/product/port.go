package product

import (
	"syloria-demo/domain"
	prdctHndlr "syloria-demo/rest/handler/product"
)

type Service interface {
	prdctHndlr.Service
}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(productID int) error
	Update(p domain.Product) (*domain.Product, error)
}
