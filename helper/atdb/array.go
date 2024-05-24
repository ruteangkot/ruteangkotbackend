package atdb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddDocToArray[T any](collection *mongo.Collection, ObjectID primitive.ObjectID, arrayname string, newDoc T) (result *mongo.UpdateResult, err error) {
	filter := bson.M{"_id": ObjectID}
	update := bson.M{
		"$push": bson.M{arrayname: newDoc},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return
	}
	return
}

func DeleteDocFromArray[T any](collection *mongo.Collection, ObjectID primitive.ObjectID, arrayname string, Doc T) (result *mongo.UpdateResult, err error) {
	filter := bson.M{"_id": ObjectID}
	update := bson.M{
		"$pull": bson.M{arrayname: Doc},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return
	}
	return
}
