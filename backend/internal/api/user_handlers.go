package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sajidcodesdotcom/kira/internal/models"
	"github.com/sajidcodesdotcom/kira/internal/services"
	"github.com/sajidcodesdotcom/kira/utils"
)

type UserHandler struct {
	userRepo services.UserRepository
	validate *validator.Validate
}

func NewUserHandler(userRepo services.UserRepository, validator *validator.Validate) *UserHandler {
	return &UserHandler{userRepo: userRepo, validate: validator}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData struct {
		FullName  string `json:"full_name" validate:"required,min=2,max=100"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=8,max=100"`
		Username  string `json:"username" validate:"required,min=3,max=100"`
		AvatarURL string `json:"avatar_url" validate:"omitempty,url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "Failed to creat user (not able to read body)", http.StatusInternalServerError)
		return
	}

	if err := h.validate.Struct(userData); err != nil {
		errorString := utils.GetValidationErrors(err)
		utils.RespondWithError(w, "incorrect user Data: "+errorString, http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(userData.Password)
	if err != nil {
		utils.RespondWithError(w, "Failed to hash use password: "+err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.NewUser(userData.FullName, userData.Email, hashedPassword, userData.Username, "user", userData.AvatarURL)

	if err := h.userRepo.Create(r.Context(), user); err != nil {
		utils.RespondWithError(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("User is created: ", user.Email)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
