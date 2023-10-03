package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Product struct {
		gorm.Model
		Code  string `json:"code" validate:"required"`
		Price uint   `json:"price" validate:"required,gte=0,lte=500"`
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
	db.AutoMigrate(&Product{})
	if err != nil {
		panic("Failed to connect to db")
	}
	return db
}

func getProducts(c echo.Context) error {
	db := getDB()
	products := []Product{}
	db.Find(&products)
	return c.JSON(http.StatusOK, map[string]interface{}{"data": products})
}

func insertProduct(c echo.Context) error {
	db := getDB()
	var p Product
	if err := c.Bind(&p); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(p); err != nil {
		return err
	}
	db.Create(&p)
	return c.JSON(http.StatusCreated, p)
}

func retrieveProduct(c echo.Context) error {
	db := getDB()
	id := c.Param("id")
	var product Product
	result := db.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"data": "Record not found"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"data": product})
}

func deleteProduct(c echo.Context) error {
	db := getDB()
	id := c.Param("id")
	var product Product
	db.Delete(&product, id)
	return c.JSON(http.StatusOK, map[string]string{"message": "Deleted successfully"})
}

func updateProduct(c echo.Context) error {
	db := getDB()
	id := c.Param("id")
	var p Product
	if err := c.Bind(&p); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(p); err != nil {
		return err
	}
	db.First(&Product{}, id).Updates(&p)
	return c.JSON(http.StatusCreated, map[string]interface{}{"data": p})
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
	e.Logger.Fatal(e.Start(":1323"))
}
