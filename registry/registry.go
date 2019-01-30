package registry

import (
	"github.com/grandcolline/clean-arch-demo/adapter/controller"
	"github.com/grandcolline/clean-arch-demo/adapter/presenter"
	"github.com/grandcolline/clean-arch-demo/driver/api/handler"
	"github.com/grandcolline/clean-arch-demo/driver/datastore"
	"github.com/grandcolline/clean-arch-demo/usecase"
	"github.com/grandcolline/clean-arch-demo/usecase/inputport"
	"github.com/grandcolline/clean-arch-demo/usecase/outputport"
	"github.com/jinzhu/gorm"
)

type interactor struct {
	conn *gorm.DB
}

type Iteractor interface {
	NewAppHandler() handler.AppHandler
}

func NewInteractor(conn *gorm.DB) Iteractor {
	return &interactor{conn}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return i.NewUserHandler()
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserController())
}

func (i *interactor) NewUserController() controller.UserController {
	return controller.NewUserController(i.NewUserService())
}

func (i *interactor) NewUserService() usecase.UserService {
	return usecase.NewUserService(i.NewUserRepository(), i.NewUserPresenter())
}

func (i *interactor) NewUserRepository() inputport.UserInputport {
	return datastore.NewUserRepository(i.conn)
}

func (i *interactor) NewUserPresenter() outputport.UserOutputport {
	return presenter.NewUserPresenter()
}
