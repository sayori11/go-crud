package view

type UserCreate struct {
	Username string `json:"username" validate:"required" gorm:"unique"`
	Password string `json:"password,omitempty" validate:"required"`
}
