package main

import (
	"fmt"
	"log"

	"github.com/ArjunMalhotra07/gorm_recruiter/application"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Started main function here")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	if err := application.StartServer(); err != nil {
		log.Fatal(err)
	}
}
