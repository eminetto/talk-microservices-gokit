package vote

import "github.com/google/uuid"

type Service interface {
	Store(v Vote) (uuid.UUID, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}
func (s *service) Store(v Vote) (uuid.UUID, error) {
	//@TODO create store rules, using databases or something else
	return v.ID, nil
}
