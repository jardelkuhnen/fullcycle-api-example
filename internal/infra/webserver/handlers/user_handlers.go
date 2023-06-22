package handlers

import (
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"github.com/jardelkuhnen/api-example/internal/dto"
	"github.com/jardelkuhnen/api-example/internal/entity"
	"github.com/jardelkuhnen/api-example/internal/infra/database"
	"net/http"
	"time"
)

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{UserDB: userDB, Jwt: jwt, JwtExpiresIn: jwtExpiresIn}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var jwt dto.JwtInput
	err := json.NewDecoder(r.Body).Decode(&jwt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.UserDB.FindByEmail(jwt.Email)
	if err != nil {
		w.Write([]byte("User not found"))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if !user.ValidatePassword(jwt.Password) {
		w.Write([]byte("Invalid password"))
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})
	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
