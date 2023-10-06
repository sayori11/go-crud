//lint:file-ignore U1000 Wait
package main

import (
	"net/http"

	_ "server/docs"
	"server/handler"
	"server/repository"
	"server/service"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Products API
// @version 1.0
// @description Test API

// @host localhost:1323
// @BasePath /
// @schemes http

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

type jwtCustomClaims struct {
	id       int
	username string
	jwt.RegisteredClaims
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
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
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	p.Use(echojwt.WithConfig(config))
	p.GET("/:id", productHandler.RetrieveProduct)
	p.GET("", productHandler.GetProducts)
	p.POST("", productHandler.InsertProduct)
	p.DELETE("/:id", productHandler.DeleteProduct)
	p.PUT("/:id", productHandler.UpdateProduct)

	u := e.Group("/")
	userSvc := service.NewUserService(pgRepo)
	userHandler := handler.NewUserHandler(userSvc)
	u.POST("register", userHandler.Register)
	u.POST("login", userHandler.Login)

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	e.Logger.Fatal(e.Start(":1323"))
}
