package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/sajidcodesdotcom/kira/internal/api"
	"github.com/sajidcodesdotcom/kira/internal/services"
	"github.com/sajidcodesdotcom/kira/pkg/database"
	"github.com/sajidcodesdotcom/kira/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env variables: %v", err)
	}
	dbpool, err := database.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to DB...")
	defer dbpool.Close()

	port := utils.GetEnvOrDefault("PORT", ":8080")

	mux := http.NewServeMux()

	userRepo := services.NewPgUserRepository(dbpool.Pool)

	userHandlers := api.NewUserHandler(userRepo, validator.New())

	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			userHandlers.ListUsers(w, r)
		}
	})

	mux.HandleFunc("/api/user/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			userHandlers.CreateUser(w, r)
		}
	})

	mux.HandleFunc("/api/user/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			userHandlers.UpdateUser(w, r)
		}
	})

	mux.HandleFunc("/api/user/by-email", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			userHandlers.GetUserByEmail(w, r)
		}
	})

	mux.HandleFunc("/api/user/by-username", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			userHandlers.GetByUsername(w, r)
		}
	})

	fmt.Print("server is running now...")
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Failed to ListenAndServe: %v", err)
	}
}
