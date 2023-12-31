package handler

import (
	"errors"
	"net/http"
	"server/helpers"
	"server/model"
	"server/service"
	"server/view"
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
// @Success 200 {object} view.DataResponse[[]model.Product]
// @Failure 400 {object} view.ErrorResponse
// @Router /products [get]
func (h *ProductHandler) GetProducts(c echo.Context) error {
	products, err := h.svc.GetProducts()
	if err != nil {
		return helpers.ErrorWrap(err, http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, view.DataResponse[[]model.Product]{Data: products})
}

// @Summary Insert a product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body view.ProductCreate true "Product Body"
// @Success 201 {object} view.DataResponse[model.Product]
// @Failure 400 {object} view.ErrorResponse
// @Router /products [post]
func (h *ProductHandler) InsertProduct(c echo.Context) error {
	p := model.Product{}
	if err := c.Bind(&p); err != nil {
		return helpers.ErrorWrap(err, http.StatusBadRequest)
	}
	if err := c.Validate(p); err != nil {
		return helpers.ErrorWrap(err, http.StatusUnprocessableEntity)
	}
	product, err := h.svc.InsertProduct(p)
	if err != nil {
		return helpers.ErrorWrap(err, http.StatusBadRequest)
	}
	return c.JSON(http.StatusCreated, view.DataResponse[model.Product]{Data: product})
}

// @Summary Retrieve a product
// @Tags Products
// @Param id path int true "Product ID"
// @Accept */*
// @Produce json
// @Success 200 {object} view.DataResponse[model.Product]
// @Failure 400 {object} view.ErrorResponse
// @Router /products/{id} [get]
func (h *ProductHandler) RetrieveProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helpers.ErrorWrap(errors.New("id should be an integer"), http.StatusBadRequest)
	}
	product, err := h.svc.RetrieveProduct(int(id))
	if err != nil {
		return helpers.ErrorWrap(err, http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, view.DataResponse[model.Product]{Data: product})
}

// @Summary Delete a product
// @Tags Products
// @Param id path int true "Product ID"
// @Accept */*
// @Produce json
// @Success 200 {object} view.DataResponse[string]
// @Failure 400 {object} view.ErrorResponse
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helpers.ErrorWrap(errors.New("id should be an integer"), http.StatusBadRequest)
	}
	if err := h.svc.DeleteProduct(id); err != nil {
		return helpers.ErrorWrap(err, http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, view.DataResponse[string]{Data: "Deleted successfully"})
}

// @Summary Update a product
// @Tags Products
// @Param id path int true "Product ID"
// @Accept json
// @Produce json
// @Param product body view.ProductCreate true "product body"
// @Success 200 {object} view.DataResponse[string]
// @Failure 400 {object} view.ErrorResponse
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helpers.ErrorWrap(errors.New("id should be an integer"), http.StatusBadRequest)
	}
	p := model.Product{}
	if err := c.Bind(&p); err != nil {
		return helpers.ErrorWrap(err, http.StatusBadRequest)
	}
	if err := c.Validate(p); err != nil {
		return helpers.ErrorWrap(err, http.StatusUnprocessableEntity)
	}

	if err := h.svc.UpdateProduct(id, p); err != nil {
		return helpers.ErrorWrap(err, http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, view.DataResponse[string]{Data: "Updated successfully"})
}
