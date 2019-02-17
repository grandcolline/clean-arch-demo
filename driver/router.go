package driver

import (
	"github.com/grandcolline/clean-arch-demo/adapter/controllers"
	"github.com/grandcolline/clean-arch-demo/driver/mysql"

	"github.com/go-chi/chi"
	"net/http"
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
