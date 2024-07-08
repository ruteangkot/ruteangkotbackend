package controller

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/at"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)
func generateResetToken() (string, error) {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    return hex.EncodeToString(b), nil
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	// Generate reset token
	resetToken, err := generateResetToken()
	if err != nil {
		http.Error(w, "Failed to generate reset token", http.StatusInternalServerError)
		return
	}

	resetTokenExpiry := time.Now().Add(1 * time.Hour)

	// Update user in database
	collection := config.DB.Collection("users")
	filter := bson.M{"email": email}
	update := bson.M{
		"$set": bson.M{
			"resetToken":       resetToken,
			"resetTokenExpiry": resetTokenExpiry,
		},
	}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to update user", http.StatusInternalServerError)
		}
		return
	}

	// TODO: Send reset email with the token to the user
	// sendResetEmail(email, resetToken)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Reset token sent to your email"}`))
}
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
	var loginRequest model.LoginRequest
	_ = json.NewDecoder(r.Body).Decode(&loginRequest)

	collection := config.DB.Collection("users")
	var user model.User
	err := collection.FindOne(context.Background(), bson.M{"email": loginRequest.Email}).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
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