package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	user, err := NewUser("Jardel", "j@.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jardel", user.Name)
	assert.Equal(t, "j@.com", user.Email)
}

func TestUser_ValidatePasswork(t *testing.T) {
	user, err := NewUser("Jardel", "j@.com", "123456")

	assert.Nil(t, err)
	assert.NotEqual(t, "123456", user.Password)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
}
