package model

type (
	TokenResponse struct {
		Token string `json:"token"`
	}
	DataResponse[T Product | []Product | User | string] struct {
		Data T `json:"data"`
	}
	ErrorResponse struct {
		Error string `json:"error"`
	}
)
