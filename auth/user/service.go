package user

import "fmt"

type Service interface {
	ValidateUser(email, password string) error
}

type service struct{}

func NewService() Service {
	return &service{}
}
func (s *service) ValidateUser(email, password string) error {
	//@TODO create validation rules, using databases or something else
	if email == "eminetto@gmail.com" && password != "1234567" {
		return fmt.Errorf("Invalid user")
	}
	return nil
}
