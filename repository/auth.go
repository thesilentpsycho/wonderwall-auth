package repository

import (
	"bitbucket.org/libertywireless/wonderwall-auth/models"
)

type AuthenticationEngine interface {
	Login(userID string, password string) models.JWTToken
	RefreshToken(token models.JWTToken) models.JWTToken
	InvalidateToken(token models.JWTToken) error
}

type authenticator struct {
}

func (a *authenticator) Login(userID string, password string) models.JWTToken {

}

func (a *authenticator) RefreshToken(token models.JWTToken) models.JWTToken {

}

func (a *authenticator) InvalidateToken(token models.JWTToken) error {

}
