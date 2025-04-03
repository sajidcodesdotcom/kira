package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/sajidcodesdotcom/kira/internal/api"
	"github.com/sajidcodesdotcom/kira/internal/middleware"
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

	router := http.NewServeMux()

	userRepo := services.NewPgUserRepository(dbpool.Pool)

	userHandlers := api.NewUserHandler(userRepo, validator.New())
	authHandlers := api.NewAuthHandler(userRepo, validator.New())

	router.Handle("/api/users", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.ListUsers)))

	router.HandleFunc("POST /api/auth/login", authHandlers.Login)
	router.HandleFunc("POST /api/auth/register", authHandlers.Register)

	router.Handle("PUT /api/user/update", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.UpdateUser)))

	router.Handle("GET /api/user/by-email", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.GetUserByEmail)))

	router.Handle("GET /api/user/by-username", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.GetByUsername)))

	fmt.Print("server is running now...")
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to ListenAndServe: %v", err)
	}
}
