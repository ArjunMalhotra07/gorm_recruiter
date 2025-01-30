package config

import (
	"os"

	"gorm.io/gorm"
)

type Config struct {
	MySql MySQL
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
	return &Config{MySql: mySQL}
}
