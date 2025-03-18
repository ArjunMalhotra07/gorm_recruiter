package config

import (
	"fmt"
	"log"
	"os"

	pb "github.com/ArjunMalhotra07/gorm_recruiter/proto"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Config struct {
	MySql         MySQL
	Microservices Microservices
}
type Microservices struct {
	EmailService pb.EmailServiceClient
}

type MySQL struct {
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDBName   string
	Driver        *gorm.DB
}

func NewConfig(fileName string) *Config {
	mySQL := MySQL{}
	//! MySQL
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		panic("Env vars DBuser not set")
	}
	mySQL.MysqlUser = dbUser
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		panic("Env vars dbPassword not set")
	}
	mySQL.MysqlPassword = dbPassword
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		panic("Env vars dbName not set")
	}
	mySQL.MysqlDBName = dbName
	dbHOST := os.Getenv("DB_HOST")
	if dbHOST == "" {
		panic("Env vars dbHOST not set")
	}
	mySQL.MysqlHost = dbHOST
	dbPORT := os.Getenv("DB_PORT")
	if dbPORT == "" {
		panic("Env vars dbPORT not set")
	}
	mySQL.MysqlPort = dbPORT

	microservices := &Microservices{}
	//! Email
	emailDomain := os.Getenv("EMAIL_HOST")
	if emailDomain == "" {
		panic("Env vars EMAIL_HOST not set")
	}
	emailPort := os.Getenv("EMAIL_PORT")
	if emailPort == "" {
		panic("Env vars EMAIL_PORT not set")
	}
	emailConnectionString := fmt.Sprintf("%s:%s", emailDomain, emailPort)
	conn, err := grpc.Dial(emailConnectionString, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to email service: %v", err)
	}
	fmt.Println("Connected to email service at ", emailConnectionString)
	emailClient := pb.NewEmailServiceClient(conn)
	microservices.EmailService = emailClient
	return &Config{MySql: mySQL, Microservices: *microservices}
}
