package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sajidcodesdotcom/kira/pkg/database"
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
}
