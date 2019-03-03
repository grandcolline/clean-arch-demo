package driver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/grandcolline/clean-arch-demo/adapter/controller"
	"github.com/grandcolline/clean-arch-demo/adapter/gateway"
	"github.com/grandcolline/clean-arch-demo/driver/mysql"
)

// Serve サーバのルーティング設定
func Serve() {
	r := chi.NewRouter()
	r.Mount("/users", userRouter())
	http.ListenAndServe(":8080", r)
}

// userRouter ユーザ用のサブルーター
func userRouter() http.Handler {
	// ユーザゲートウェイの作成
	conn := mysql.Connect()
	userGateway := gateway.NewUserGateway(conn)

	// ユーザコントローラの作成
	logger := &Logger{}
	userController := controller.NewUserController(userGateway, logger)

	// ルーティング
	r := chi.NewRouter()
	r.HandleFunc("/", userController.FindAll)
	r.HandleFunc("/{userID}", userController.FindByID)

	return r
}
