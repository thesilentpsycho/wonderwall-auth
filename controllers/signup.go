package controllers

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/libertywireless/wonderwall-auth/models"
	"github.com/sirupsen/logrus"
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
		logrus.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v := validator.New()
	err = v.Struct(form)
	json.NewEncoder(w).Encode(err.(validator.ValidationErrors))

}

func (c *authController) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yo Bhuwania"))
}

func NewAuthController() AuthController {
	return &authController{}
}
