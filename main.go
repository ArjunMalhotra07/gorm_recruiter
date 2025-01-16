package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	driver, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
		return
	}
	fmt.Println("Database connected successfully:", driver)
	// Auto-migrate all models
	if err := driver.AutoMigrate(&models.User{}, &models.Job{}, &models.Resume{}, &models.Education{}, &models.Experience{}, &models.JobApplication{}); err != nil {
		log.Fatalf("failed to auto-migrate database: %v", err)
		return
	}
	var app *models.App = application.New(driver)
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
