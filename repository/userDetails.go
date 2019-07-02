package repository

import (
	"bitbucket.org/libertywireless/wonderwall-auth/models"
)

type M map[string]interface{}

type UserEngine interface {
	Create(details models.UserDetails) error
	UpdateMultiple(userid string, fieldValues M) error
	UpdateSingle(userid string, field string, newValue interface{}) error
}

type userRepo struct {
}

func (r *userRepo) Create(details models.UserDetails) error {
	return nil
	//To do
}

func (r *userRepo) UpdateMultiple(userid string, fieldValues M) error {
	return nil
	//To do
}

func (r *userRepo) UpdateSingle(userid string, field string, newValue interface{}) error {
	return nil
	//To do
}
