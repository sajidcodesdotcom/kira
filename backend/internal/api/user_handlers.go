package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

type UserData struct {
	ID        uuid.UUID `json:"id" validate:"required,uuid"`
	FullName  string    `json:"full_name" validate:"required,min=2,max=100"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8,max=100"`
	Username  string    `json:"username" validate:"required,min=3,max=100"`
	AvatarURL string    `json:"avatar_url" validate:"omitempty,url"`
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	r = r.WithContext(ctx)
	defer cancel()

	var userEmail struct {
		Email string `json:"email" validate:"required,email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userEmail); err != nil {
		utils.RespondWithError(w, "Error reading body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(userEmail); err != nil {
		errorString := utils.GetValidationErrors(err)
		utils.RespondWithError(w, errorString, http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetByEmail(r.Context(), userEmail.Email)
	if err != nil {
		utils.RespondWithError(w, "Failed to get User by email: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, user, http.StatusOK)
}

func (h *UserHandler) GetByUsername(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	r = r.WithContext(ctx)
	defer cancel()

	var Username struct {
		Username string `json:"username" validate:"required,min=3,max=100"`
	}

	if err := json.NewDecoder(r.Body).Decode(&Username); err != nil {
		utils.RespondWithError(w, "Error reading body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(Username); err != nil {
		errorString := utils.GetValidationErrors(err)
		utils.RespondWithError(w, errorString, http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetByUsername(r.Context(), Username.Username)
	if err != nil {
		utils.RespondWithError(w, "Failed to get User by usernamee: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, user, http.StatusOK)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	r = r.WithContext(ctx)
	defer cancel()

	var userData UserData

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		utils.RespondWithError(w, "Error reading body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(userData); err != nil {
		errorString := utils.GetValidationErrors(err)
		utils.RespondWithError(w, errorString, http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(userData.Password)
	if err != nil {
		utils.RespondWithError(w, "Failed to hash password: "+err.Error(), http.StatusInternalServerError)
	}

	if err != nil {
		utils.RespondWithError(w, "Failed to parse UUID: "+err.Error(), http.StatusInternalServerError)
	}

	existingUser, err := h.userRepo.GetByID(r.Context(), userData.ID)
	if err != nil {
		utils.RespondWithError(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	existingUser.Email = userData.Email
	existingUser.FullName = userData.FullName
	existingUser.AvatarURL = userData.AvatarURL
	existingUser.Password = hashedPassword

	if err := h.userRepo.Update(r.Context(), existingUser); err != nil {
		utils.RespondWithError(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
	}

	utils.RespondWithJSON(w, existingUser, http.StatusOK)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	r = r.WithContext(ctx)
	defer cancel()

	var usersData []UserData

	users, err := h.userRepo.List(r.Context(), 5, 0)
	if err != nil {
		utils.RespondWithError(w, "Failed to get user list: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for _, user := range users {
		userData := UserData{
			ID:        user.ID,
			FullName:  user.FullName,
			Email:     user.Email,
			Username:  user.Username,
			AvatarURL: user.AvatarURL,
		}

		usersData = append(usersData, userData)
	}

	utils.RespondWithJSON(w, usersData, http.StatusOK)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	r = r.WithContext(ctx)
	userID := r.URL.Query().Get("id")
	id, err := uuid.Parse(userID)
	if err != nil {
		utils.RespondWithError(w, "Failed to parse UUID: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.userRepo.Delete(r.Context(), id); err != nil {
		utils.RespondWithError(w, "Failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer cancel()
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	r = r.WithContext(ctx)
	defer cancel()

	userID := r.Context().Value("user_id").(uuid.UUID)

	user, err := h.userRepo.GetByID(r.Context(), userID)
	if err != nil {
		utils.RespondWithError(w, "Failed to get user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, user, http.StatusOK)
}
