package repository

import (
	"errors"
	"server/model"

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

func (repo *PGRepository) ValidateUser(user model.UserCreate) error {
	db := repo.DB
	userDB := model.User{}
	if result := db.First(&userDB, "username = ?", user.Username); result.Error != nil {
		return errors.New("incorrect username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
		return errors.New("incorrect password")
	}

	return nil
}