package repository

import (
	"context"

	datastore "bitbucket.org/libertywireless/wonderwall-auth/infra"
	models "bitbucket.org/libertywireless/wonderwall-auth/models/db"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type M map[string]interface{}

type UserEngine interface {
	Create(details *models.UserDetails) error
	UpdateMultiple(userid string, fieldValues M) error
	UpdateSingle(userid string, field string, newValue interface{}) error
}

type userRepo struct {
	Datastore  *datastore.MongoDatastore
	Collection *mongo.Collection
}

func (r *userRepo) Create(details *models.UserDetails) error {
	_, err := r.Collection.InsertOne(context.Background(), details)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) UpdateMultiple(userid string, fieldValues M) error {
	return nil
	//To do
}

func (r *userRepo) UpdateSingle(userid string, field string, newValue interface{}) error {
	return nil
	//To do
}

func NewUserRepository(store *datastore.MongoDatastore, collection string) *userRepo {
	if collection == "" {
		//Ye sb sahi nahi ho rha hai
		return nil
	}
	return &userRepo{
		Collection: store.DB.Collection(collection),
		Datastore:  store}
}
