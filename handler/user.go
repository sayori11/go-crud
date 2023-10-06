package handler

import (
	"net/http"
	"server/helpers"
	"server/model"
	"server/service"
	"server/view"

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
// @Param user body view.UserCreate true "User Body"
// @Success 201 {object} view.DataResponse[model.User]
// @Failure 400 {object} view.ErrorResponse
// @Router /register [post]
func (h *UserHandler) Register(c echo.Context) error {
	u := model.User{}
	if err := c.Bind(&u); err != nil {
		return helpers.ErrorWrap(err, http.StatusBadRequest)
	}
	if err := c.Validate(u); err != nil {
		return helpers.ErrorWrap(err, http.StatusUnprocessableEntity)
	}
	user, err := h.svc.Register(u)
	if err != nil {
		return helpers.ErrorWrap(err, http.StatusBadRequest)
	}
	return c.JSON(http.StatusCreated, view.DataResponse[model.User]{Data: user})
}

// @Summary Login
// @Tags User
// @Accept json
// @Produce json
// @Param user body view.UserCreate true "User Body"
// @Success 200 {object} view.TokenResponse
// @Failure 400 {object} view.ErrorResponse
// @Router /login [post]
func (h *UserHandler) Login(c echo.Context) error {
	u := view.UserCreate{}
	if err := c.Bind(&u); err != nil {
		return helpers.ErrorWrap(err, http.StatusBadRequest)
	}
	if err := c.Validate(u); err != nil {
		return helpers.ErrorWrap(err, http.StatusUnprocessableEntity)
	}

	token, err := h.svc.Login(u)
	if err != nil {
		return helpers.ErrorWrap(err, http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, view.TokenResponse{Token: token})
}
