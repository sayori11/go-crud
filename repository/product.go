package repository

import "server/model"

func (repo *PGRepository) GetProducts() ([]model.Product, error) {
	db := repo.DB
	products := []model.Product{}
	db.Find(&products)
	return products, nil
}

func (repo *PGRepository) InsertProduct(product model.Product) (model.Product, error) {
	db := repo.DB
	db.Create(&product)
	return product, nil
}

func (repo *PGRepository) UpdateProduct(id int, product model.Product) error {
	db := repo.DB
	result := db.First(&model.Product{}, id).Updates(&product)
	return result.Error
}

func (repo *PGRepository) RetrieveProduct(id int) (model.Product, error) {
	db := repo.DB
	product := model.Product{}
	result := db.First(&product, id)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return product, nil
}

func (repo *PGRepository) DeleteProduct(id int) error {
	db := repo.DB
	result := db.Delete(&model.Product{}, id)
	return result.Error
}
