package user

import (
	"auth/security"
	"errors"
)

type Service interface {
	ValidateUser(email, password string) (string, error)
	ValidateToken(token string) (string, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) ValidateUser(email, password string) (string, error) {
	//@TODO create validation rules, using databases or something else
	if email == "eminetto@gmail.com" && password != "1234567" {
		return "nil", errors.New("Invalid user")
	}
	token, err := security.NewToken(email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *service) ValidateToken(token string) (string, error) {
	t, err := security.ParseToken(token)
	if err != nil {
		return "", err
	}
	tData, err := security.GetClaims(t)
	if err != nil {
		return "", err
	}
	return tData["email"].(string), nil
}
