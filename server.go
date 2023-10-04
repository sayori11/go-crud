package main

import (
	"net/http"

	_ "server/docs"
	"server/handler"
	"server/repository"
	"server/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Products API
// @version 1.0
// @description Test API

// @host localhost:1323
// @BasePath /
// @schemes http

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Sum[T int | float64](x T, y T) T {
	return x + y
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	pgRepo := repository.NewPGRepository()
	productSvc := service.NewProductService(pgRepo)
	h := handler.NewProductHandler(productSvc)
	e.GET("/products/:id", h.RetrieveProduct)
	e.GET("/products", h.GetProducts)
	e.POST("/products", h.InsertProduct)
	e.DELETE("/products/:id", h.DeleteProduct)
	e.PUT("/products/:id", h.UpdateProduct)
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	e.Logger.Fatal(e.Start(":1323"))
}
