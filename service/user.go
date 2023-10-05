package service

import (
	"server/model"
	"server/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.IRepository
}

type jwtCustomClaims struct {
	username string
	jwt.RegisteredClaims
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

func (svc UserService) Login(user model.UserCreate) (string, error) {
	err := svc.repo.ValidateUser(user)
	if err != nil {
		return "", err
	}
	claims := &jwtCustomClaims{
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}