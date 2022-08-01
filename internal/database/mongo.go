package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func formatObjectId(hex string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return objectID, err
	}
	return objectID, nil
}

func formatObjectIdMultiple(hex []string) ([]primitive.ObjectID, error) {
	var list []primitive.ObjectID

	oids := make([]primitive.ObjectID, len(hex))
	for _, i := range hex {
		objectId, err := primitive.ObjectIDFromHex(i)
		if err != nil {
			return nil, err
		}
		oids = append(oids, objectId)
	}
	return list, nil
}

type MongoUser struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string              `json:"first_name" bson:"firstName"`
	LastName  string              `json:"last_name" bson:"lastName"`
	Email     string              `json:"email" bson:"email"`
}

func (mongo *Mongo) GetUser(id string) (*MongoUser, error) {
	objectID, err := formatObjectId(id)
	if err != nil {
		fmt.Println("inside get FCM")
		return nil, err
	}

	filter := bson.D{{"_id", objectID}}
	coll := mongo.Con.Database("slizzr").Collection("users")
	result, err := coll.FindOne(context.TODO(), filter).DecodeBytes()
	var output MongoUser
	bson.Unmarshal(result, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

func (mongo *Mongo) GetMultipleUser(ids []string) ([]*MongoUser, error) {

	objectIDs, err := formatObjectIdMultiple(ids)
	if err != nil {
		return nil, err
	}

	query := bson.M{"_id": bson.M{"$in": objectIDs}}

	coll := mongo.Con.Database("slizzr").Collection("users")
	cursor, err := coll.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer func() {
		cursor.Close(context.Background())
	}()
	var output []*MongoUser
	for cursor.Next(context.Background()) {
		var temp *MongoUser
		cursor.Decode(&temp)
		output = append(output, temp)
	}

	return output, nil
}
