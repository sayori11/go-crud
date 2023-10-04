package service

import "server/model"

type IService interface {
	GetProducts() ([]model.Product, error)
	InsertProduct(model.Product) (model.Product, error)
	UpdateProduct(int, model.Product) error
	RetrieveProduct(int) (model.Product, error)
	DeleteProduct(int) error
}
