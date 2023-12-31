package repository

import (
	"server/model"
	"server/view"
)

type IRepository interface {
	GetProducts() ([]model.Product, error)
	InsertProduct(model.Product) (model.Product, error)
	UpdateProduct(int, model.Product) error
	RetrieveProduct(int) (model.Product, error)
	DeleteProduct(int) error
	CreateUser(model.User) (model.User, error)
	ValidateUser(view.UserCreate) (model.User, error)
}
