package controller

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
)

func Getdatarouteangkot(respw http.ResponseWriter, req *http.Request) {
	resp, _:= atdb.GetAllDoc[[]model.RuteAngkot](config.Mongoconn, "data json", bson.M{})
	helper.WriteJSON(respw, http.StatusOK, resp)
	
}