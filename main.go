package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/application"
	"github.com/ArjunMalhotra07/gorm_recruiter/bootstrap"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Started main function")
	dsn := "root:example@tcp(127.0.0.1:3307)/crud_db?charset=utf8mb4&parseTime=True&loc=Local"
	driver, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
