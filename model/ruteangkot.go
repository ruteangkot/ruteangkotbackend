package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RuteAngkot struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Rute           string             `bson:"Rute,omitempty" json:"Rute,omitempty"`
	JamOperasional string             `bson:"Jam Operasional,omitempty" json:"Jam Operasional,omitempty"`
	Tarif          string             `bson:"Tarif,omitempty" json:"Tarif,omitempty"`
}

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
}
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
}