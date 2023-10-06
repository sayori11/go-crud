package view

import "server/model"

type (
	TokenResponse struct {
		Token string `json:"token"`
	}
	DataResponse[T model.Product | []model.Product | model.User | string] struct {
		Data T `json:"data"`
	}
	ErrorResponse struct {
		Message string `json:"message"`
	}
)
