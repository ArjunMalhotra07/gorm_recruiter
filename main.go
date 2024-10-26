package main

import (
	"fmt"
	"log"

	"github.com/ArjunMalhotra07/gorm_recruiter/application"
	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Started main function")
	dsn := "root:Witcher_Arjun7@tcp(127.0.0.1:3306)/New_DB?charset=utf8mb4&parseTime=True&loc=Local"
	driver, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
		return
	}
	fmt.Println("Database connected successfully:", driver)
	var app *models.App = application.New(driver)
	if err := application.StartServer(app); err != nil {
		log.Fatal(err)
	}
}
