package handler

import (
	"context"

	"net/http"

	"github.com/grandcolline/clean-arch-demo/adapter/controller"
	"github.com/grandcolline/clean-arch-demo/entity"
	"github.com/labstack/echo"
)

type userHandler struct {
	userController controller.UserController
}

type UserHandler interface {
	CreateUser(c echo.Context) error
	GetUsers(c echo.Context) error
}

func NewUserHandler(uc controller.UserController) UserHandler {
	return &userHandler{userController: uc}
}

func (uh *userHandler) CreateUser(c echo.Context) error {

	// リクエスト用のEntityを作成
	req := &entity.User{}

	// bind
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.ResponseError{Message: err.Error()})
	}

	// validate
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.ResponseError{Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := uh.userController.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, "success")
}
func (uh *userHandler) GetUsers(c echo.Context) error {

	req := &entity.User{}

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.ResponseError{Message: err.Error()})
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	u, err := uh.userController.GetUsers()
	if err != nil {
		// システム内のエラー
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, u)
}
