package user_test

import (
	"auth/user"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	service := user.NewService()
	t.Run("invalid user", func(t *testing.T) {
		_, err := service.ValidateUser(context.Background(), "eminetto@gmail.com", "invalid")
		assert.NotNil(t, err)
		assert.Equal(t, user.ErrInvalidUser, err)
	})
	t.Run("valid user", func(t *testing.T) {
		token, err := service.ValidateUser(context.Background(), "eminetto@gmail.com", "1234567")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
	})
}
