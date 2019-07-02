package controllers

import (
	"net/http"

	"bitbucket.org/libertywireless/wonderwall-auth/repository"
)

type SignInController interface {
	SignIn(w http.ResponseWriter, r *http.Request)
}

type signInController struct {
	authRepository repository.AuthenticationEngine
}

func (c *signInController) SignIn(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Yo Bhuwania"))
}

func NewSignInController(authRepository repository.AuthenticationEngine) *signInController {
	return &signInController{
		authRepository: authRepository}
}
