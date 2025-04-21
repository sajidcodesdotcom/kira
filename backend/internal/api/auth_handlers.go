package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sajidcodesdotcom/kira/internal/auth"
	"github.com/sajidcodesdotcom/kira/internal/models"
	"github.com/sajidcodesdotcom/kira/internal/services"
	"github.com/sajidcodesdotcom/kira/utils"
)

type AuthHandler struct {
	userRepo services.UserRepository
	validate *validator.Validate
}

func NewAuthHandler(userRepo services.UserRepository, validate *validator.Validate) *AuthHandler {
	return &AuthHandler{
		userRepo: userRepo,
		validate: validate,
	}
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	FullName  string `json:"full_name" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required,min=3,max=100"`
	Password  string `json:"password" validate:"required,min=8,max=100"`
	AvatarURL string `json:"avatar_url" validate:"omitempty,url"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// create timeout context
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	r = r.WithContext(ctx)
	defer cancel()

	var userData RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "Failed to login user (not able to read body)", http.StatusInternalServerError)
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

	user := models.NewUser(uuid.New(), userData.FullName, userData.Email, hashedPassword, userData.Username, "user", userData.AvatarURL)

	token, err := auth.GenerateToken(user)

	if err := h.userRepo.Create(r.Context(), user); err != nil {
		utils.RespondWithError(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := AuthResponse{
		Token: token,
		User: &models.User{
			ID:        user.ID,
			FullName:  user.FullName,
			Email:     user.Email,
			Username:  user.Username,
			Role:      user.Role,
			AvatarURL: user.AvatarURL,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	auth.SetTokenCookie(w, token)

	utils.RespondWithJSON(w, response, http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// create timeout context
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	r = r.WithContext(ctx)
	defer cancel()

	var userData LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "Failed to login user (not able to read body)", http.StatusInternalServerError)
		return
	}

	if err := h.validate.Struct(userData); err != nil {
		errorString := utils.GetValidationErrors(err)
		utils.RespondWithError(w, "incorrect user Data: "+errorString, http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetByEmail(r.Context(), userData.Email)
	if err != nil {
		utils.RespondWithError(w, "Failed to find user in the DB: "+err.Error(), http.StatusUnauthorized)
		return
	}

	isUserExist := utils.CheckPassword(user.Password, userData.Password)
	if !isUserExist {
		utils.RespondWithError(w, "Incorrect Password: ", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		utils.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth.SetTokenCookie(w, token)

	response := AuthResponse{
		Token: token,
		User: &models.User{
			ID:        user.ID,
			FullName:  user.FullName,
			Email:     user.Email,
			Username:  user.Username,
			Role:      user.Role,
			AvatarURL: user.AvatarURL,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	utils.RespondWithJSON(w, response, http.StatusCreated)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	auth.ClearTokenCookie(w)

	utils.RespondWithJSON(w, "successfully logged out", http.StatusOK)
}
