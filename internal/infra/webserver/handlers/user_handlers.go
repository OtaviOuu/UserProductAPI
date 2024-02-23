package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/crudGolangAPI/internal/dto"
	"github.com/crudGolangAPI/internal/entity"
	"github.com/crudGolangAPI/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB       database.UserInterfaceDB
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userDB database.UserInterfaceDB, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       userDB,
		Jwt:          jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var userInput dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Valida email
	u, err := h.UserDB.FindByEmail(userInput.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Valida password
	if !u.ValidatePassword(userInput.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Payload do token
	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),                                        // Usuario
		"exp": time.Now().Add(time.Duration(h.JwtExpiresIn)).Unix(), // Tempo de duração do token
	})
	acessToken := struct {
		AcessToken string `json:"acess_token"`
	}{
		AcessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acessToken)

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
