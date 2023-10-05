package repository

import "server/model"

func (repo *PGRepository) CreateUser(user model.User) (model.User, error) {
	db := repo.DB
	result := db.Create(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (repo *PGRepository) ValidateUser(user model.User) (bool, error) {
	return true, nil
}
