package main

import (
	"fmt"
	// "net/http"

	// "github.com/go-chi/chi"
	"github.com/grandcolline/clean-arch-demo/config"
	"github.com/grandcolline/clean-arch-demo/driver/api/router"
	"github.com/grandcolline/clean-arch-demo/driver/api/validater"
	"github.com/grandcolline/clean-arch-demo/driver/datastore"
	"github.com/grandcolline/clean-arch-demo/registry"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	config.Init()
	conn := datastore.NewMySqlDB()

	// interactor
	r := registry.NewInteractor(conn)

	// 依存解決
	h := r.NewAppHandler()

	// echo
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// validate
	e.Validator = &validater.CustomValidator{Validator: validator.New()}

	// router
	router.NewRouter(e, h)

	// chiで書き換える
	// router := chi.NewRouter()
	// router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	// router.Get("/users", h.handler.GetUsers)

	// DB stop
	defer func() {
		if err := conn.Close(); err != nil {
			// e.Logger.Fatal(fmt.Sprintf("Failed to close: %v", err))
			fmt.Println("DB ERROR")
		}
	}()

	// server start
	e.Logger.Fatal(e.Start(":8080"))
	// http.ListenAndServe(":8080", router)
}
