package repository

import (
	datastore "bitbucket.org/libertywireless/wonderwall-auth/infra"
	models "bitbucket.org/libertywireless/wonderwall-auth/models"
	dbmodels "bitbucket.org/libertywireless/wonderwall-auth/models/db"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type AuthenticationEngine interface {
	Login(userID string, password string) models.JWTToken
	RefreshToken(token models.JWTToken) models.JWTToken
	InvalidateToken(token models.JWTToken) error
	CreateLogin(info *dbmodels.AuthInfo) error
}

type authenticator struct {
	Datastore  *datastore.MongoDatastore
	Collection *mongo.Collection
}

func (a *authenticator) CreateLogin(info *dbmodels.AuthInfo) error {
	return nil
}

func (a *authenticator) Login(userID string, password string) models.JWTToken {
	return models.JWTToken{
		Token: "wfwdffefw"}
}

func (a *authenticator) RefreshToken(token models.JWTToken) models.JWTToken {
	return models.JWTToken{
		Token: "wfwdffefw"}
}

func (a *authenticator) InvalidateToken(token models.JWTToken) error {
	return nil
}
