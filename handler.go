package main

import (
	"bitbucket.org/libertywireless/wonderwall-auth/config"
	"bitbucket.org/libertywireless/wonderwall-auth/constants"
	"bitbucket.org/libertywireless/wonderwall-auth/controllers"
	datastore "bitbucket.org/libertywireless/wonderwall-auth/infra"
	"bitbucket.org/libertywireless/wonderwall-auth/repository"
	"bitbucket.org/libertywireless/wonderwall-auth/wlog"
	"github.com/go-chi/chi"
)

func GetHandler() *chi.Mux {
	router := GetRouter()

	config := config.GetConfig()
	logger := wlog.Logger

	store := datastore.NewDatastore(*config, logger)
	datastore.PopulateIndex(store, "auth")
	authRepository := repository.NewAuthRepository(store, constants.AuthCollectionName)
	userRepository := repository.NewUserRepository(store, constants.UserDetailsCollectionName)

	signInCtrl := controllers.NewSignInController(authRepository)
	signUpCtrl := controllers.NewSignUpController(userRepository, authRepository)

	router.Post("/signup", signUpCtrl.SignUp)
	router.Post("/signin", signInCtrl.SignIn)
	return router
}
