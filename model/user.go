package model

import "gorm.io/gorm"

type (
	UserCreate struct {
		Username string `json:"username" validate:"required" gorm:"unique"`
		Password string `json:"password" validate:"required"`
	}
	User struct {
		gorm.Model
		UserCreate
	}
)
