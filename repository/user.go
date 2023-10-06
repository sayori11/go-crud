package repository

import (
	"errors"
	"server/model"
	"server/view"

	"golang.org/x/crypto/bcrypt"
)

func (repo *PGRepository) CreateUser(user model.User) (model.User, error) {
	db := repo.DB
	result := db.Create(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (repo *PGRepository) ValidateUser(user view.UserCreate) (model.User, error) {
	db := repo.DB
	userDB := model.User{}
	if result := db.First(&userDB, "username = ?", user.Username); result.Error != nil {
		return model.User{}, errors.New("incorrect username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		return model.User{}, errors.New("incorrect password")
	}

	return userDB, nil
}
