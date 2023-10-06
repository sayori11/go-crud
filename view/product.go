package view

type ProductCreate struct {
	Code  string `json:"code" validate:"required" example:"A45"`
	Price uint   `json:"price" validate:"required,gte=0,lte=500" example:"200"`
}
