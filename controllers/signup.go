package controllers

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/libertywireless/wonderwall-auth/models"
	"gopkg.in/go-playground/validator.v9"
)

type AuthController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}

type authController struct {
}

func (c *authController) SignUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var form models.SignUpForm
	err := decoder.Decode(&form)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v := validator.New()
	err = v.Struct(form)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

}

func (c *authController) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yo Bhuwania"))
}

func NewAuthController() AuthController {
	return &authController{}
}
