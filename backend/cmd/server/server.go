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

	// auth handlers
	authHandlers := api.NewAuthHandler(userRepo, validator.New())
	router.HandleFunc("POST /api/auth/login", authHandlers.Login)
	router.HandleFunc("POST /api/auth/register", authHandlers.Register)

	// user handlers
	userHandlers := api.NewUserHandler(userRepo, validator.New())
	router.Handle("/api/users", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.ListUsers)))
	router.Handle("PUT /api/user/update", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.UpdateUser)))
	router.Handle("GET /api/user/by-email", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.GetUserByEmail)))
	router.Handle("GET /api/user/by-username", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.GetByUsername)))
	router.Handle("DELETE /api/user/delete", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.Delete)))

	// project handlers
	projectRepo := services.NewPgProjectRepository(dbpool.Pool)
	projectHandlers := api.NewProjectHandler(projectRepo, validator.New())
	router.Handle("POST /api/project/create", middleware.AuthMiddleware(http.HandlerFunc(projectHandlers.CreateProject)))
	router.Handle("GET /api/projects", middleware.AuthMiddleware(http.HandlerFunc(projectHandlers.ListProjects)))
	router.Handle("GET /api/project/by-id", middleware.AuthMiddleware(http.HandlerFunc(projectHandlers.GetProjectByID)))
	router.Handle("PUT /api/project/update", middleware.AuthMiddleware(http.HandlerFunc(projectHandlers.UpdateProject)))
	router.Handle("DELETE /api/project/delete", middleware.AuthMiddleware(http.HandlerFunc(projectHandlers.DeleteProject)))
	router.Handle("GET /api/project/by-owner", middleware.AuthMiddleware(http.HandlerFunc(projectHandlers.GetProjectsByOwner)))

	fmt.Print("server is running now...")
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to ListenAndServe: %v", err)
	}
}
