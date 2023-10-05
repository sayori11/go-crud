package service

import "server/model"

type IProductService interface {
	GetProducts() ([]model.Product, error)
	InsertProduct(model.Product) (model.Product, error)
	UpdateProduct(int, model.Product) error
	RetrieveProduct(int) (model.Product, error)
	DeleteProduct(int) error
}

type IUserService interface {
	Register(model.User) (model.User, error)
	Login(model.UserCreate) error
}
