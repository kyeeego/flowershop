package config

import (
	"os"
	"strconv"
)

type Config struct {
	DbName     string
	DbHost     string
	DbPort     int
	DbUsername string
	DbPassword string

	Port int
	Prod bool
}

func Init() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	prod, err := strconv.ParseBool(os.Getenv("PROD"))
	if err != nil {
		return nil, err
	}

	return &Config{
		DbName:     os.Getenv("DB_NAME"),
		DbHost:     os.Getenv("DB_HOST"),
		DbUsername: os.Getenv("DB_USER"),
		DbPort:     dbPort,
		DbPassword: os.Getenv("DB_PASSWORD"),

		Port: port,
		Prod: prod,
	}, nil
}
