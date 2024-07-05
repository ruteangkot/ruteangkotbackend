package controller

import (
	"context"
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
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)
func RequestResetPassword(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		helper.JSON(w, http.StatusBadRequest, helper.Message("Invalid request body"))
		return
	}

	resetToken := helper.GenerateResetToken()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("users")
	filter := bson.M{"email": requestData.Email}
	update := bson.M{"$set": bson.M{"resetToken": resetToken, "resetTokenExpiry": time.Now().Add(1 * time.Hour)}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedUser model.User
	err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			helper.JSON(w, http.StatusNotFound, helper.Message("Email not found"))
			return
		}
		helper.JSON(w, http.StatusInternalServerError, helper.Message("Failed to update user"))
		return
	}

	if err := helper.SendResetPasswordEmail(updatedUser.Email, resetToken); err != nil {
		helper.JSON(w, http.StatusInternalServerError, helper.Message("Failed to send email"))
		return
	}

	helper.JSON(w, http.StatusOK, helper.Message("Reset password email sent"))
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		helper.JSON(w, http.StatusBadRequest, helper.Message("Invalid request body"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.JSON(w, http.StatusInternalServerError, helper.Message("Failed to hash password"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("users")
	filter := bson.M{"resetToken": requestData.Token, "resetTokenExpiry": bson.M{"$gt": time.Now()}}
	update := bson.M{"$set": bson.M{"password": string(hashedPassword)}, "$unset": bson.M{"resetToken": "", "resetTokenExpiry": ""}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil || result.ModifiedCount == 0 {
		helper.JSON(w, http.StatusBadRequest, helper.Message("Invalid or expired token"))
		return
	}

	helper.JSON(w, http.StatusOK, helper.Message("Password reset successfully"))
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