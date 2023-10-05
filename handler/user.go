package handler

import (
	"net/http"
	"server/model"
	"server/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	svc service.IUserService
}

func NewUserHandler(svc service.IUserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// @Summary Register
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.UserCreate true "User Body"
// @Success 201 {object} model.DataResponse[User]
// @Failure 400 {object} model.ErrorResponse
// @Router /register [post]
func (h *UserHandler) Register(c echo.Context) error {
	u := model.User{}
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Binding error"})
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}
	user, err := h.svc.Register(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusCreated, model.DataResponse[model.User]{Data: user})
}

// @Summary Login
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.UserCreate true "User Body"
// @Success 200 {object} model.TokenResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /login [post]
func (h *UserHandler) Login(c echo.Context) error {
	u := model.UserCreate{}
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Binding error"})
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	token, err := h.svc.Login(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, model.TokenResponse{Token: token})
}
