package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Slizzr/slizzr-user-service/internal/config"
	"github.com/Slizzr/slizzr-user-service/internal/database"
	"github.com/Slizzr/slizzr-user-service/internal/http"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// api documentation notation
	// @title API documentation
	// @version 1.0.0
	// @host localhost:3000
	// @BasePath /

	// load .env into os.env
	_ = godotenv.Load(".env")
	conf, err := config.NewFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.MongoConfig.Uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to db")
	}

	mongo := database.Mongo{
		Con: client,
	}

	//start server
	server := http.Server{
		Config:  conf,
		Context: context.Background(),
		Mongo:   &mongo,
	}

	fmt.Printf("starting server on port: http://localhost:%s \n", server.Config.Server.Port)
	err = server.ListenAndServe(fmt.Sprintf(":%s", server.Config.Server.Port))
	if err != nil {
		panic(err)
	}
}
