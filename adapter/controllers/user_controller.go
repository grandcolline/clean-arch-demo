package controllers

import (
	"fmt"
	"github.com/grandcolline/clean-arch-demo/adapter/gateway"
	"github.com/grandcolline/clean-arch-demo/adapter/interfaces"
	"github.com/grandcolline/clean-arch-demo/entity"
	"github.com/grandcolline/clean-arch-demo/usecase"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(conn *gorm.DB, logger interfaces.Logger) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &gateway.UserRepository{
				Conn: conn,
			},
			Logger: logger,
		},
	}
}

func (controller *UserController) Create(c interfaces.Context) {
	type (
		Request struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		Response struct {
			UserID int `json:"user_id"`
		}
	)
	req := Request{}
	c.Bind(&req)
	user := entity.User{Name: req.Name, Email: req.Email}

	id, err := controller.Interactor.Add(user)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "user_controller: cannot add user"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	res := Response{UserID: id}
	c.JSON(201, res)
}

// func (controller *UserController) FindByName(c interfaces.Context) {
func (controller *UserController) FindByName(w http.ResponseWriter, r *http.Request) {
	type (
		Response struct {
			UserID int `json:"user_id"`
		}
	)
	name := r.URL.Query().Get("name")

	users, err := controller.Interactor.FindByName(name)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "user_controller: cannot fond user"))
		http.Error(w, err.Error(), 500)
		return
	}
	// res := Response{UserID: users[0].ID}
	// c.JSON(201, res)
	w.Write([]byte(fmt.Sprintf("id:%d", users[0].ID)))
}
