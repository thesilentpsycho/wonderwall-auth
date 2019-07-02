package models

type SignUpForm struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	EmailID   string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}
