package route

import (
	"log"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func URL(w http.ResponseWriter, r *http.Request) {
	SetEnv()

	var method, path string = r.Method, r.URL.Path
	switch {
	case method == "GET" && path == "/":
		controller.GetHome(w, r)
	case method == "GET" && path == "/refresh/token":
		controller.GetNewToken(w, r)
	case method == "GET" && path == "/data/user":
		controller.GetDataUser(w, r)
	case method == "POST" && helper.URLParam(path, "/webhook/nomor/:nomorwa"):
		controller.PostInboxNomor(w, r)
	default:
		controller.NotFound(w, r)
	}
}

func SetEnv() {
	if config.ErrorMongoconn != nil {
		log.Println(config.ErrorMongoconn.Error())
	}
	profile, err := atdb.GetOneDoc[model.Profile](config.Mongoconn, "profile", primitive.M{})
	if err != nil {
		log.Println(err)
	}
	config.PublicKeyWhatsAuth = profile.PublicKey
}
