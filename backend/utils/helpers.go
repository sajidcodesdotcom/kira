package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func CheckPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

type Error struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, message string, statusCode int) {
	RespondWithJSON(w, Error{Error: message}, statusCode)
}

func RespondWithJSON(w http.ResponseWriter, message any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, "Failed to Encode and response messsage "+err.Error(), http.StatusInternalServerError)
	}
}

func GetEnvOrDefault(key, defaultValue string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return defaultValue
	}
	return value
}

func GetValidationErrors(err error) string {
	var errMsgs []string

	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			errMsgs = append(errMsgs, e.Field()+" is required")
		case "email":
			errMsgs = append(errMsgs, e.Field()+" must be a valid email")
		case "min":
			errMsgs = append(errMsgs, e.Field()+" must be at least "+e.Param()+" characters")
		case "max":
			errMsgs = append(errMsgs, e.Field()+" must not exceed "+e.Param()+" characters")
		case "url":
			errMsgs = append(errMsgs, e.Field()+" msut be a valid email")
		default:
			errMsgs = append(errMsgs, e.Error())
		}
	}

	return strings.Join(errMsgs, "; ")
}
