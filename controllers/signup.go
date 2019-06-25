package controllers

import "net/http"

type AuthController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}

type authController struct {
}

func (c *authController) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yo Bhuwania"))
}

func (c *authController) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yo Bhuwania"))
}

func NewAuthController() AuthController {
	return &authController{}
}
