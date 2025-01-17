package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ArjunMalhotra07/gorm_recruiter/application"
	"github.com/ArjunMalhotra07/gorm_recruiter/bootstrap"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Started main function")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbName)
	fmt.Println(dsn)

	var db *gorm.DB
	retryAttempts := 5
	retryInterval := 5 * time.Second

	for i := 0; i < retryAttempts; i++ {
		driver, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			db = driver
			fmt.Println("Database connected successfully:", db)
			break
		}
		log.Printf("Failed to connect to the database (attempt %d/%d): %v", i+1, retryAttempts, err)
		time.Sleep(retryInterval)
	}

	if db == nil {
		log.Fatalf("Failed to connect to the database after %d attempts", retryAttempts)
		return
	}

	// Auto-migrate all models
	if err := db.AutoMigrate(&models.User{}, &models.Job{}, &models.Resume{}, &models.Education{}, &models.Experience{}, &models.JobApplication{}); err != nil {
		log.Fatalf("failed to auto-migrate database: %v", err)
		return
	}
	var app *models.App = application.New(db)
	bootstrap.RegisterMetrics()
	go func() {
		// Serve Prometheus metrics at /metrics endpoint
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Prometheus metrics available at /metrics")
		log.Fatal(http.ListenAndServe(":9100", nil)) // Prometheus server runs on port 9100
	}()
	if err := application.StartServer(app); err != nil {
		log.Fatal(err)
	}
}