package config

import (
	"errors"
	"fmt"
	"os"
)

func NewFromEnvironment() (Config, error) {
	mongo, err := NewMongoConfig()
	if err != nil {
		return Config{}, err
	}
	return Config{
		Server: Server{
			Port:    os.Getenv("APP_PORT"),
			SiteURL: fmt.Sprintf("%s:%s", os.Getenv("APP_URL"), os.Getenv("APP_PORT")),
		},
		MongoConfig: mongo,
	}, nil
}

type MongoConfig struct {
	Uri string `json:"MONGO_URI"`
}

func NewMongoConfig() (MongoConfig, error) {
	output := MongoConfig{
		Uri: os.Getenv("MONGO_CONNECTION_URI"),
	}

	if output.Uri == "" {
		return MongoConfig{}, errors.New("connection uri missing in env")
	}
	return output, nil
}

type Config struct {
	Server      Server
	MongoConfig MongoConfig
}

type Server struct {
	Port    string
	SiteURL string
}
