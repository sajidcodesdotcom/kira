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
		if r.Method == "POST" {
			userHandlers.CreateUser(w, r)
		}

		if r.Method == "GET" {
			w.Write([]byte("teh gest request is working fine"))
		}
	})

	fmt.Print("server is running now...")
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Failed to ListenAndServe: %v", err)
	}
}
