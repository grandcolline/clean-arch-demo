package driver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/grandcolline/clean-arch-demo/adapter/controller"
	"github.com/grandcolline/clean-arch-demo/adapter/gateway"
	"github.com/grandcolline/clean-arch-demo/driver/config"
	"github.com/grandcolline/clean-arch-demo/driver/mysql"
)

// conf アプリケーション設定
var conf config.AppConf

// Serve はサーバのルーティング設定を行います
func Serve() {
	// 設定の読み込み
	conf.Init()

	r := chi.NewRouter()
	r.Mount("/users", userRouter())
	http.ListenAndServe(":"+conf.Port, r)
}

/*
userRouter はユーザ用のサブルーターとして、
`/users/*`のルーティングを行います。
*/
func userRouter() http.Handler {
	// ユーザゲートウェイの作成
	conn := mysql.Connect()
	userGateway := gateway.NewUserGateway(conn)

	// ユーザコントローラの作成
	logger := &Logger{}
	userController := controller.NewUserController(userGateway, logger)

	// ルーティング
	r := chi.NewRouter()
	r.Get("/", userController.FindAll) // GET /users
	r.Post("/", userController.Add)    // POST /users
	r.Route("/{userID:[0-9-]+}", func(r chi.Router) {
		r.Get("/", userController.FindByID)  // GET /users/{userID}
		r.Put("/", userController.Change)    // PUT /users/{userID}
		r.Delete("/", userController.Delete) // DELETE /users/{userID}
	})

	return r
}
