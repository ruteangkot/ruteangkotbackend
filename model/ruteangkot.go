package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type RuteAngkot struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Rute           string             `bson:"Rute,omitempty" json:"Rute,omitempty"`
	JamOperasional string             `bson:"Jam Operasional,omitempty" json:"Jam Operasional,omitempty"`
	Tarif          string             `bson:"Tarif,omitempty" json:"Tarif,omitempty"`
}