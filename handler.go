package main

import (
	"bitbucket.org/libertywireless/wonderwall-auth/controllers"
	"github.com/go-chi/chi"
)

func GetHandler() *chi.Mux {
	router := GetRouter()

	authController := controllers.NewAuthController()
	router.Post("/signup", authController.SignUp)
	router.Post("/signin", authController.SignIn)
	return router
}
