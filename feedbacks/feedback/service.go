package feedback

import "github.com/google/uuid"

type Service interface {
	Store(f Feedback) (uuid.UUID, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) Store(f Feedback) (uuid.UUID, error) {
	//@TODO create store rules, using databases or something else
	return f.ID, nil
}
