package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sajidcodesdotcom/kira/internal/models"
	"github.com/sajidcodesdotcom/kira/internal/services"
	"github.com/sajidcodesdotcom/kira/pkg/database"
)

func main() {
	ctx := context.Background()
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

	userRepo := services.NewPgUserRepository(dbpool.Pool)
	if err != nil {
		log.Fatal(err)
	}
	user := models.NewUser("sajid", "sdfdsfsdf@gmail.com", "sdfdf", "dsfdsfdsfs", "sdfsdfk", "sdfsdfdsf")
	err = userRepo.Create(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

}
