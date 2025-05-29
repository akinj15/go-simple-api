package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	// Create a new user
	user, err := NewUser("John Doe", "jon.doe@email.com", "password123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "jon.doe@email.com", user.Email)
}

func TestValidadePassword(t *testing.T) {
	// Create a new user
	user, err := NewUser("John Doe", "jon.doe@email.com", "password123")

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("password123"))
	assert.False(t, user.ValidatePassword("password123a"))

}
