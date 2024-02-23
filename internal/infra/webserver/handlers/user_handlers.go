package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/crudGolangAPI/internal/dto"
	"github.com/crudGolangAPI/internal/entity"
	"github.com/crudGolangAPI/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterfaceDB
}

func NewUserHandler(userDB database.UserInterfaceDB) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
