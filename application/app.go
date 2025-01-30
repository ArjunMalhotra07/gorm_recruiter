package application

import (
	"log"
	"net/http"

	"github.com/ArjunMalhotra07/gorm_recruiter/bootstrap"
	"github.com/ArjunMalhotra07/gorm_recruiter/internal/config"
	"github.com/ArjunMalhotra07/gorm_recruiter/internal/db"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/ArjunMalhotra07/gorm_recruiter/routes"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewApp(config *config.Config) *models.App {
	return &models.App{Router: routes.AppRoutes(config), Config: config}
}

func StartServer() error {
	cfg := config.NewConfig(".env")
	sqlDB, err := db.NewMySQLDb(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	if err := db.Migrate(sqlDB.DB); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
	cfg.MySql.Driver = sqlDB.DB
	app := NewApp(cfg)
	bootstrap.RegisterMetrics()
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Prometheus metrics available at /metrics")
		log.Fatal(http.ListenAndServe(":9100", nil)) // Prometheus server runs on port 9100
	}()
	server := &http.Server{Addr: ":8080", Handler: app.Router}
	err = http.ListenAndServe(server.Addr, server.Handler)
	/*
		or
		err := http.ListenAndServe(":8080", app.router)
	*/
	if err != nil {
		return err
	}
	return nil
}
