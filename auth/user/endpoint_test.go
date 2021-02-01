package user

import (
	"context"
	"testing"
)

func TestMakeValidateUserEndpoint(t *testing.T) {
	s := NewService()
	endpoint := makeValidateUserEndpoint(s)
	t.Run("valid user", func(t *testing.T) {
		req := validateUserRequest{
			Email:    "eminetto@gmail.com",
			Password: "1234567",
		}
		_, err := endpoint(context.Background(), req)
		if err != nil {
			t.Errorf("expected %v received %v", nil, err)
		}
	})
	t.Run("invalid user", func(t *testing.T) {
		req := validateUserRequest{
			Email:    "eminetto@gmail.com",
			Password: "123456",
		}
		_, err := endpoint(context.Background(), req)
		if err == nil {
			t.Errorf("expected %v received %v", ErrInvalidUser, err)
		}
	})
}
