package db

import (
	"log"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadConfig() *Config {
	defaultDBHost := "go_db"
	defaultDBPort := "5432"
	defaultDBUser := "postgres"
	defaultDBPassword := "1234"
	defaultDBName := "postgres"

	Host := os.Getenv("DB_HOST")

	if Host == "" {
		log.Println("DB_HOST environment variable not set, setting to default: " + defaultDBHost)
		Host = defaultDBHost
	}

	Port := os.Getenv("DB_PORT")

	if Port == "" {
		log.Println("DB_PORT environment variable not set, setting to default: " + defaultDBPort)
		Port = defaultDBPort
	}

	User := os.Getenv("DB_USER")

	if User == "" {
		log.Println("DB_USER environment variable not set, setting to default: " + defaultDBUser)
		User = defaultDBUser
	}

	Password := os.Getenv("DB_PASSWORD")

	if Password == "" {
		log.Println("DB_PASSWORD environment variable not set, setting to default: " + defaultDBPassword)
		Password = defaultDBPassword
	}

	DBName := os.Getenv("DB_NAME")

	if DBName == "" {
		log.Println("DB_NAME environment variable not set, setting to default: " + defaultDBName)
		DBName = defaultDBName
	}

	return &Config{
		Host:     Host,
		Port:     Port,
		User:     User,
		Password: Password,
		DBName:   DBName,
	}
}
