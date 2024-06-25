package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gocroot/helper"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

var mongoinfo = model.DBInfo{
	DBString: MongoString,
	DBName:   "Angkutankotabdg",
}

var Mongoconn, ErrorMongoconn = helper.MongoConnect(mongoinfo)

var DB *mongo.Database

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Angkutankotabdg:naikangkot001@cluster02.hqmgzeb.mongodb.net/"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("userDB")
	log.Println("Connected to MongoDB!")
}