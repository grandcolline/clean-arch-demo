package driver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/grandcolline/clean-arch-demo/adapter/controllers"
	"github.com/grandcolline/clean-arch-demo/driver/mysql"
)

func Serve() {
	logger := &Logger{}
	conn := mysql.Connect()
	userController := controllers.NewUserController(conn, logger)

	r := chi.NewRouter()
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		userController.FindByName(w, r)
	})
	http.ListenAndServe(":8080", r)
}
