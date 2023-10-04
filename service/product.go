package service

import (
	"errors"
	"server/model"
	"server/repository"
)

type ProductService struct {
	repo repository.IRepository
}

func NewProductService(repo repository.IRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (svc *ProductService) GetProducts() ([]model.Product, error) {
	return svc.repo.GetProducts()
}

func (svc *ProductService) InsertProduct(product model.Product) (model.Product, error) {
	return svc.repo.InsertProduct(product)
}

func (svc *ProductService) UpdateProduct(id int, product model.Product) error {
	if id < 0 {
		return errors.New("ID is invalid")
	}
	return svc.repo.UpdateProduct(id, product)
}

func (svc *ProductService) RetrieveProduct(id int) (model.Product, error) {
	if id < 0 {
		return model.Product{}, errors.New("ID is invalid")
	}
	return svc.repo.RetrieveProduct(id)
}

func (svc *ProductService) DeleteProduct(id int) error {
	if id < 0 {
		return errors.New("ID is invalid")
	}
	return svc.repo.DeleteProduct(id)
}
