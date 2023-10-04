package main

import (
	"net/http"

	_ "server/docs"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Products API
// @version 1.0
// @description Test API

// @host localhost:1323
// @BasePath /
// @schemes http

type (
	ProductCreate struct {
		Code  string `json:"code" validate:"required" example:"A45"`
		Price uint   `json:"price" validate:"required,gte=0,lte=500" example:"200"`
	}
	Product struct {
		gorm.Model
		ProductCreate
	}
	DataResponse[T Product | []Product | string] struct {
		Data T `json:"data"`
	}
	ErrorResponse struct {
		Error string `json:"error"`
	}
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

func getDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}
	db.AutoMigrate(&Product{})
	return db
}

// @Summary Get a list of products
// @Tags Products
// @Accept */*
// @Produce json
// @Success 200 {object} DataResponse[[]Product]
// @Router /products [get]
func getProducts(c echo.Context) error {
	db := getDB()
	products := []Product{}
	db.Find(&products)
	return c.JSON(http.StatusOK, DataResponse[[]Product]{Data: products})
}

// @Summary Insert a product
// @Tags Products
// @Accept json
// @Produce json
// @Param object body ProductCreate true "Product Body"
// @Success 200 {object} DataResponse[Product]
// @Failure 400 {object} ErrorResponse
// @Router /products [post]
func insertProduct(c echo.Context) error {
	db := getDB()
	var p Product
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Binding error"})
	}
	if err := c.Validate(p); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
	}
	db.Create(&p)
	return c.JSON(http.StatusCreated, DataResponse[Product]{Data: p})
}

// @Summary Retrieve a product
// @Tags Products
// @Param id path int true "Product ID"
// @Accept */*
// @Produce json
// @Success 200 {object} DataResponse[Product]
// @Failure 400 {object} ErrorResponse
// @Router /products/{id} [get]
func retrieveProduct(c echo.Context) error {
	db := getDB()
	id := c.Param("id")
	var product Product
	result := db.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Product not found"})
	}
	return c.JSON(http.StatusOK, DataResponse[Product]{Data: product})
}

// @Summary Delete a product
// @Tags Products
// @Param id path int true "Product ID"
// @Accept */*
// @Produce json
// @Success 200 {object} DataResponse[string]
// @Router /products/{id} [delete]
func deleteProduct(c echo.Context) error {
	db := getDB()
	id := c.Param("id")
	var product Product
	db.Delete(&product, id)
	return c.JSON(http.StatusOK, DataResponse[string]{Data: "Deleted successfully"})
}

// @Summary Update a product
// @Tags Products
// @Param id path int true "Product ID"
// @Accept json
// @Produce json
// @Param object body ProductCreate true "product body"
// @Success 200 {object} DataResponse[Product]
// @Failure 400 {object} ErrorResponse
// @Router /products/{id} [put]
func updateProduct(c echo.Context) error {
	db := getDB()
	id := c.Param("id")
	var p Product
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Binding error"})
	}
	if err := c.Validate(p); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
	}
	db.First(&Product{}, id).Updates(&p)
	return c.JSON(http.StatusOK, DataResponse[Product]{Data: p})
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
	e.GET("/products/:id", retrieveProduct)
	e.GET("/products", getProducts)
	e.POST("/products", insertProduct)
	e.DELETE("/products/:id", deleteProduct)
	e.PUT("/products/:id", updateProduct)
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	e.Logger.Fatal(e.Start(":1323"))
}
