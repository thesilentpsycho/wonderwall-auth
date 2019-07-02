package repository

import (
	"context"
	"errors"

	datastore "bitbucket.org/libertywireless/wonderwall-auth/infra"
	models "bitbucket.org/libertywireless/wonderwall-auth/models"
	dbmodels "bitbucket.org/libertywireless/wonderwall-auth/models/db"
	"bitbucket.org/libertywireless/wonderwall-auth/wlog"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type AuthenticationEngine interface {
	Login(userID string, password string) models.JWTToken
	RefreshToken(token models.JWTToken) models.JWTToken
	InvalidateToken(token models.JWTToken) error
	CreateLogin(info *dbmodels.AuthInfo) (string, error)
}

type authenticator struct {
	Datastore  *datastore.MongoDatastore
	Collection *mongo.Collection
}

func (a *authenticator) CreateLogin(info *dbmodels.AuthInfo) (string, error) {
	wlog.Logger.Infoln("Creating new Login for :", info.EmailID)

	res, err := a.Collection.InsertOne(context.Background(), info)

	if err != nil {
		wlog.Logger.Errorln("Could not insert into collection :", a.Collection.Name())

		if err, ok := err.(mongo.WriteErrors); ok {
			if (err)[0].Code == 11000 {
				return "", errors.New("Already Exists")
			}
		}
		return "", errors.New("Internal server error")
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (a *authenticator) Login(userID string, password string) models.JWTToken {
	return models.JWTToken{
		Token: "wfwdffefw"}
}

func (a *authenticator) RefreshToken(token models.JWTToken) models.JWTToken {
	return models.JWTToken{
		Token: "yuitiut"}
}

func (a *authenticator) InvalidateToken(token models.JWTToken) error {
	return nil
}

func NewAuthRepository(store *datastore.MongoDatastore, collection string) *authenticator {
	return &authenticator{
		Collection: store.DB.Collection(collection),
		Datastore:  store}
}
