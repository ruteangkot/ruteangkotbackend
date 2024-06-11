package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/at"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
)

func Getdatarouteangkot(respw http.ResponseWriter, req *http.Request) {
	resp, _:= atdb.GetAllDoc[[]model.RuteAngkot](config.Mongoconn, "data json", bson.M{})
	helper.WriteJSON(respw, http.StatusOK, resp)
	
}

func CreateRoute(respw http.ResponseWriter, req *http.Request) {
	resp, _:= atdb.GetAllDoc[[]model.RuteAngkot](config.Mongoconn, "data json", bson.M{})
	helper.WriteJSON(respw, http.StatusOK, resp)
	
}

func UpdateRoute(respw http.ResponseWriter, req *http.Request) {
	var rute model.RuteAngkot
	err := json.NewDecoder(req.Body).Decode(&rute)
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
	}
	dt, err:= atdb.GetOneDoc[model.RuteAngkot](config.Mongoconn,"data json",bson.M{"_id":rute.ID})
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
	}
	dt.JamOperasional = rute.JamOperasional
	dt.Rute = rute.Rute
	dt.Tarif = rute.Tarif
	_, err= atdb.ReplaceOneDoc(config.Mongoconn,"data json",bson.M{"_id":rute.ID},dt)
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
	}
	helper.WriteJSON(respw, http.StatusOK, dt)
	
}

func DeleteRoute(respw http.ResponseWriter, req *http.Request) {
	resp, _:= atdb.GetAllDoc[[]model.RuteAngkot](config.Mongoconn, "data json", bson.M{})
	helper.WriteJSON(respw, http.StatusOK, resp)
	
}