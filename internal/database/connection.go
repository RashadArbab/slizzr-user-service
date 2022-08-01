package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	Con *mongo.Client
}
