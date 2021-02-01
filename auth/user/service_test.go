package user_test

import (
	"auth/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	service := user.NewService()
	t.Run("invalid user", func(t *testing.T) {
		_, err := service.ValidateUser("eminetto@gmail.com", "invalid")
		assert.NotNil(t, err)
		assert.Equal(t, "Invalid user", err.Error())
	})
	t.Run("valid user", func(t *testing.T) {
		token, err := service.ValidateUser("eminetto@gmail.com", "1234567")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
	})
}
