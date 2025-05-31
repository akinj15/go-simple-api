package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/akinj15/go-api/configs"
	"github.com/akinj15/go-api/internal/dto"
	"github.com/akinj15/go-api/internal/entity"
	"github.com/akinj15/go-api/internal/infra/database"
	"github.com/go-chi/jwtauth/v5"
)

type UserHandler struct {
	UserDB database.UserInterface
	Cfg    *configs.Config
	Jwt    *jwtauth.JWTAuth
}

func NewUserHandler(db database.UserInterface, cfg *configs.Config) *UserHandler {

	return &UserHandler{
		UserDB: db,
		Cfg:    cfg,
		Jwt:    cfg.TokenAuth,
	}
}

func (h *UserHandler) CreateJWT(w http.ResponseWriter, r *http.Request) {
	var input dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// tokenAuth = h.Jwt.New("HS256", []byte("secret"), nil) // replace with secret key

	user, err := h.UserDB.FindByEmail(input.Email)
	if err != nil || user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if user.Name != input.Name {
		http.Error(w, "Invalid user name", http.StatusUnauthorized)
		return
	}

	if !user.ValidatePassword(input.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	_, token, _ := h.Jwt.Encode(map[string]any{
		"sub": user.ID.String(),
	})
	log.Default().Println("User authenticated successfully:", token)

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	// Implementation for creating a user
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
