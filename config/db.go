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





