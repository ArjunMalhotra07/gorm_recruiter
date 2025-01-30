package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/application"
	"github.com/ArjunMalhotra07/gorm_recruiter/bootstrap"
	"github.com/ArjunMalhotra07/gorm_recruiter/pkg/config"
	"github.com/ArjunMalhotra07/gorm_recruiter/pkg/db"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Println("Started main function")
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Load configuration
	cfg := config.NewConfig(".env")
	// Initialize database
	sqlDB, err := db.NewMySQLDb(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	// Run database migrations
	if err := db.Migrate(sqlDB.DB); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
	// Initialize application
	app := application.New(sqlDB.DB)
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
