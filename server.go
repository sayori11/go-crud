package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string `json:"code"`
	Price uint   `json:"price" gorm:"default: 400"`
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
	db.First(&Product{}, id).Updates(&p)
	return c.JSON(http.StatusCreated, map[string]interface{}{"data": p})
}

func main() {
	e := echo.New()
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
