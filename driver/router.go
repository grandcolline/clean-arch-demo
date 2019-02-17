package driver

import (
	"github.com/grandcolline/clean-arch-demo/adapter/controllers"
	"github.com/grandcolline/clean-arch-demo/driver/mysql"

	// "github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"net/http"
)

// var Router *gin.Engine

func Serve() {
	logger := &Logger{}
	conn := mysql.Connect()
	userController := controllers.NewUserController(conn, logger)

	// router := gin.Default()
	// router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	// router.GET("/users", func(c *gin.Context) { userController.FindByName(c) })
	// Router = router

	r := chi.NewRouter()
	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		userController.FindByName(w, r)
	})
	http.ListenAndServe(":8080", r)
}
