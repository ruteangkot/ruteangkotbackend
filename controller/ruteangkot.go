package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/at"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()

	collection := config.DB.Collection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var credentials model.User
	_ = json.NewDecoder(r.Body).Decode(&credentials)

	collection := config.DB.Collection("users")
	var user model.User
	err := collection.FindOne(context.Background(), bson.M{"username": credentials.Username}).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// TODO: Generate JWT token and send it to the user

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Login successful")
}
func Getdatarouteangkot(respw http.ResponseWriter, req *http.Request) {
	resp, _:= atdb.GetAllDoc[[]model.RuteAngkot](config.Mongoconn, "data json", bson.M{})
	helper.WriteJSON(respw, http.StatusOK, resp)
	
}

func CreateRoute(respw http.ResponseWriter, req *http.Request) {
	var rute model.RuteAngkot
	err := json.NewDecoder(req.Body).Decode(&rute)
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	_, err = atdb.InsertOneDoc(config.Mongoconn,"data json",rute)
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	ruteangkots, err  := atdb.GetAllDoc[[]model.RuteAngkot](config.Mongoconn,"data json",bson.M{})
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	helper.WriteJSON(respw, http.StatusOK, ruteangkots)
	
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
	var rute model.RuteAngkot
	err := json.NewDecoder(req.Body).Decode(&rute)
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	err = atdb.DeleteOneDoc(config.Mongoconn,"data json",bson.M{"_id":rute.ID})
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	ruteangkot, err  := atdb.GetAllDoc[[]model.RuteAngkot](config.Mongoconn,"data json",bson.M{"_id":rute.ID})
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	helper.WriteJSON(respw, http.StatusOK, ruteangkot)

	
}