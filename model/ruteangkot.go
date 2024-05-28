package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type RuteAngkot struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Rute           string             `bson:"rute" json:"Rute"`
	JamOperasional string             `bson:"jam_operasional" json:"Jam Operasional"`
	Tarif          string             `bson:"tarif" json:"Tarif"`
}