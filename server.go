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

	p := e.Group("/products")
	pgRepo := repository.NewPGRepository()
	productSvc := service.NewProductService(pgRepo)
	productHandler := handler.NewProductHandler(productSvc)
	p.GET("/:id", productHandler.RetrieveProduct)
	p.GET("", productHandler.GetProducts)
	p.POST("", productHandler.InsertProduct)
	p.DELETE("/:id", productHandler.DeleteProduct)
	p.PUT("/:id", productHandler.UpdateProduct)

	u := e.Group("/")
	userSvc := service.NewUserService(pgRepo)
	userHandler := handler.NewUserHandler(userSvc)
	u.POST("register", userHandler.Register)

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	e.Logger.Fatal(e.Start(":1323"))
}
