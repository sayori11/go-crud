package service

import (
	"server/model"
	"server/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.IRepository
}

func NewUserService(repo repository.IRepository) *UserService {
	return &UserService{repo: repo}
}

func (svc UserService) Register(user model.User) (model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return model.User{}, err
	}
	user.Password = string(hash)
	return svc.repo.CreateUser(user)
}

func (svc UserService) Login(user model.UserCreate) error {
	return nil
}
