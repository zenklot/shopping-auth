package app

import (
	"shopping/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/users/", userController.FindAll)
	router.POST("/api/users/signup", userController.Create)
	router.POST("/api/users/signin", userController.Sigin)

	return router
}
