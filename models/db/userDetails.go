package models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type UserDetails struct {
	RawId            primitive.ObjectID `bson:"_id"`
	FirstName        string             `bson:"first_name" validate:"required"`
	LastName         string             `bson:"last_name"`
	EmailID          string             `bson:"email" validate:"required,email"`
	PrimaryContact   string             `bson:"primary_contact"`
	SecondaryContact string             `bson:"secondary_contact"`
	City             string             `bson:"city"`
}
