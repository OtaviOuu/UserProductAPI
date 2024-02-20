package database

import (
	"testing"

	"github.com/crudGolangAPI/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		t.Error(err)
	}

	user, _ := entity.NewUser("joao", "joao@gmail.com", "senhajoao")
	userDB := NewUserDB(db)

	err = userDB.Create(user)
	assert.Nil(err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(err)

	assert.Equal(user.ID, userFound.ID)
	assert.Equal(user.Name, userFound.Name)
	assert.Equal(user.Email, userFound.Email)
	assert.NotNil(userFound.Password)

}

func TestFindByEmail(t *testing.T) {
	assert := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("joao", "joao@gmail.com", "senhajoao")
	userDB := NewUserDB(db)

	err = userDB.Create(user)
	assert.Nil(err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(err)
	assert.Equal(user.ID, userFound.ID)
	assert.Equal(user.Email, userFound.Email)
	assert.Equal(user.Name, userFound.Name)
	assert.NotNil(userFound.Password)

}
