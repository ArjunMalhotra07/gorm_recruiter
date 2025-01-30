package db

import (
	"fmt"
	"log"
	"time"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"github.com/ArjunMalhotra07/gorm_recruiter/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDb struct to hold the database instance
type MySQLDb struct {
	DB *gorm.DB
}

// NewMySQLDb initializes and connects to MySQL database
func NewMySQLDb(cfg *config.Config) (*MySQLDb, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySql.MysqlUser, cfg.MySql.MysqlPassword, cfg.MySql.MysqlHost,
		cfg.MySql.MysqlPort, cfg.MySql.MysqlDBName,
	)

	retryAttempts := 5
	retryInterval := 5 * time.Second
	sqlDB := &MySQLDb{}

	for i := 0; i < retryAttempts; i++ {
		driver, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB.DB = driver
			fmt.Println("Database connected successfully")
			return sqlDB, nil
		}
		log.Printf("Failed to connect to the database (attempt %d/%d): %v", i+1, retryAttempts, err)
		time.Sleep(retryInterval)
	}

	return nil, fmt.Errorf("failed to connect to the database after %d attempts", retryAttempts)
}

// Migrate runs the auto-migration for all models
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}, &models.Job{}, &models.Resume{},
		&models.Education{}, &models.Experience{}, &models.JobApplication{}); err != nil {
		log.Fatalf("failed to auto-migrate database: %v", err)
		return err
	}
	return nil
}
