package config

import (
	"context"
	"log"
	"time"

	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var IPPort, Net = helper.GetAddress()

var PhoneNumber string

func SetEnv() {
	if ErrorMongoconn != nil {
		log.Println(ErrorMongoconn.Error())
	}
	profile, err := atdb.GetOneDoc[model.Profile](Mongoconn, "profile", primitive.M{})
	if err != nil {
		log.Println(err)
	}
	PublicKeyWhatsAuth = profile.PublicKey
	WAAPIToken = profile.Token
}
var client *mongo.Client

// ConnectDB initializes the database connection and creates the unique index.
func ConnectDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI("mongodb+srv://Angkutankotabdg:naikangkot001@cluster02.hqmgzeb.mongodb.net/")
    var err error
    client, err = mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Create unique index on email
    createUniqueIndex("userDB", "users", "email")
}

// GetDBCollection returns a collection from the connected database.
func GetDBCollection(collectionName string) *mongo.Collection {
    return client.Database("userDB").Collection(collectionName)
}

// createUniqueIndex ensures a unique index on the specified field.
func createUniqueIndex(database, collection, field string) {
    coll := client.Database(database).Collection(collection)
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    indexModel := mongo.IndexModel{
        Keys: bson.D{{Key: email, Value: 1}}, // index in ascending order
        Options: options.Index().SetUnique(true),
    }
    _, err := coll.Indexes().CreateOne(ctx, indexModel)
    if err != nil {
        log.Fatalf("Failed to create unique index on %s: %v", field, err)
    }
}