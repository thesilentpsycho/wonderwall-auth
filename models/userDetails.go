package models

type UserDetails struct {
	FirstName        string `bson:"first_name" validate:"required"`
	LastName         string `bson:"last_name"`
	EmailID          string `bson:"email" validate:"required,email"`
	PrimaryContact   string `bson:"primary_contact"`
	SecondaryContact string `bson:"secondary_contact"`
	City             string `bson:"city"`
}
