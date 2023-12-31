package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" validate:"required" gorm:"unique"`
	Password string `json:"password,omitempty" validate:"required"`
}
