package handler

import (
	"net/http"
	"server/model"
	"server/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	ProductHandler struct {
		svc service.IProductService
	}
)

func NewProductHandler(svc service.IProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

// @Summary Get a list of products
// @Tags Products
// @Accept */*
// @Produce json
// @Success 200 {object} model.DataResponse[[]Product]
// @Failure 400 {object} model.ErrorResponse
// @Router /products [get]
func (h *ProductHandler) GetProducts(c echo.Context) error {
	products, err := h.svc.GetProducts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, model.DataResponse[[]model.Product]{Data: products})
}

// @Summary Insert a product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body model.ProductCreate true "Product Body"
// @Success 201 {object} model.DataResponse[Product]
// @Failure 400 {object} model.ErrorResponse
// @Router /products [post]
func (h *ProductHandler) InsertProduct(c echo.Context) error {
	p := model.Product{}
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Binding error"})
	}
	if err := c.Validate(p); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}
	product, err := h.svc.InsertProduct(p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusCreated, model.DataResponse[model.Product]{Data: product})
}

// @Summary Retrieve a product
// @Tags Products
// @Param id path int true "Product ID"
// @Accept */*
// @Produce json
// @Success 200 {object} model.DataResponse[Product]
// @Failure 400 {object} model.ErrorResponse
// @Router /products/{id} [get]
func (h *ProductHandler) RetrieveProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "ID should be an integer"})
	}
	product, err := h.svc.RetrieveProduct(int(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, model.DataResponse[model.Product]{Data: product})
}

// @Summary Delete a product
// @Tags Products
// @Param id path int true "Product ID"
// @Accept */*
// @Produce json
// @Success 200 {object} model.DataResponse[string]
// @Failure 400 {object} model.ErrorResponse
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "ID should be an integer"})
	}
	if err := h.svc.DeleteProduct(id); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, model.DataResponse[string]{Data: "Deleted successfully"})
}

// @Summary Update a product
// @Tags Products
// @Param id path int true "Product ID"
// @Accept json
// @Produce json
// @Param product body model.ProductCreate true "product body"
// @Success 200 {object} model.DataResponse[string]
// @Failure 400 {object} model.ErrorResponse
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "ID should be an integer"})
	}
	p := model.Product{}
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Binding error"})
	}
	if err := c.Validate(p); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	if err := h.svc.UpdateProduct(id, p); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, model.DataResponse[string]{Data: "Updated successfully"})
}
