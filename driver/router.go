package driver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/grandcolline/clean-arch-demo/adapter/controller"
	"github.com/grandcolline/clean-arch-demo/driver/mysql"
)

// Serve サーバ設定
// ルーティング（コントローラの指定）もここで行う
func Serve() {
	logger := &Logger{}
	conn := mysql.Connect()

	r := chi.NewRouter()
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		userController := controller.NewUserController(w, conn, logger)
		// userController.FindByName(w, r)
		userController.FindAll(w, r)
	})
	r.Get("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userController := controller.NewUserController(w, conn, logger)
		// userController.FindByName(w, r)
		userController.FindByID(w, r)
	})
	http.ListenAndServe(":8080", r)
}
