package models

type AuthInfo struct {
	EmailID  string `bson:"email" validate:"required"`
	Password string `bson:"password" validate:"required"`
}
