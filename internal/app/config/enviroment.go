package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	Protocol string
	Host     string
	Port     string
	MongoUri string
}

func LoadEnv() *Environment {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	protocol := os.Getenv("GO_PROTOCOL")
	host := os.Getenv("GO_HOST")
	port := os.Getenv("GO_PORT")
	mongoUri := os.Getenv("GO_MONGO_URI")
	return &Environment{
		Host:     host,
		Port:     port,
		Protocol: protocol,
		MongoUri: mongoUri,
	}
}
