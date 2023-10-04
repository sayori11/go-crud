package model

type (
	DataResponse[T Product | []Product | string] struct {
		Data T `json:"data"`
	}
	ErrorResponse struct {
		Error string `json:"error"`
	}
)
