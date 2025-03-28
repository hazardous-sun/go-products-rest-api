package db

import "os"

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadConfig() *Config {
	Host := os.Getenv("DB_HOST")

	if Host == "" {
		panic("DB_HOST environment variable not set")
	}

	Port := os.Getenv("DB_PORT")

	if Port == "" {
		panic("DB_PORT environment variable not set")
	}

	User := os.Getenv("DB_USER")

	if User == "" {
		panic("DB_USER environment variable not set")
	}

	Password := os.Getenv("DB_PASSWORD")

	if Password == "" {
		panic("DB_PASSWORD environment variable not set")
	}

	DBName := os.Getenv("DB_NAME")

	if DBName == "" {
		panic("DB_NAME environment variable not set")
	}

	return &Config{
		Host:     Host,
		Port:     Port,
		User:     User,
		Password: Password,
		DBName:   DBName,
	}
}
