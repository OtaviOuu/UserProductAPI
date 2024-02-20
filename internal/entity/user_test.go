package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testa apenas a criação de um novo usuário -> Todos os campos são prenchidos apenas
func TestNewUser(t *testing.T) {
	assert := assert.New(t)

	user, err := NewUser("jorge", "jorge@gmail", "senhajorge123")

	assert.Nil(err)
	assert.NotEmpty(user.Password)
	assert.NotEmpty(user.ID)
	assert.Equal("jorge", user.Name)
	assert.Equal("jorge@gmail", user.Email)
}

func TestNewUser_ValidadePassword(t *testing.T) {
	assert := assert.New(t)

	user, err := NewUser("jorge", "jorge@gmail.com", "senhajorge")

	assert.Nil(err)

	assert.True(user.ValidatePassword("senhajorge"))
	assert.False(user.ValidatePassword("senha errada"))
	assert.NotEqual("senhajorge", user.Password) // Valida se há hash
}
