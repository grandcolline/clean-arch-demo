package driver

import (
	"github.com/grandcolline/clean-arch-demo/adapter/controllers"
	"github.com/grandcolline/clean-arch-demo/driver/mysql"

	"github.com/gin-gonic/gin"
	// "net/http"
	// "github.com/go-chi/chi"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	logger := &Logger{}
	conn := mysql.Connect()
	userController := controllers.NewUserController(conn, logger)

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.FindByName(c) })
	Router = router

	// r := chi.NewRouter()
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	// http.ListenAndServe(":8080", r)
}
