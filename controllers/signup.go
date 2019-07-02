package controllers

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/libertywireless/wonderwall-auth/repository"

	dbmodels "bitbucket.org/libertywireless/wonderwall-auth/models/db"
	models "bitbucket.org/libertywireless/wonderwall-auth/models/external"
	"gopkg.in/go-playground/validator.v9"
)

type SignUpController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
}

type signUpController struct {
	detailsRepo repository.UserEngine
	authRepo    repository.AuthenticationEngine
}

func (c *signUpController) SignUp(w http.ResponseWriter, r *http.Request) {
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

	authData := &dbmodels.AuthInfo{
		EmailID:  form.EmailID,
		Password: form.Password}

	err = c.authRepo.CreateLogin(authData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	userDetails := &dbmodels.UserDetails{
		EmailID:   form.EmailID,
		FirstName: form.FirstName,
		LastName:  form.LastName}

	err = c.detailsRepo.Create(userDetails)
}

func NewSignUpController(detailsRepo repository.UserEngine, authRepo repository.AuthenticationEngine) *signUpController {
	return &signUpController{
		authRepo:    authRepo,
		detailsRepo: detailsRepo}
}
