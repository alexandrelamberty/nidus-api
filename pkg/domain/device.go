package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name"`
	Mac          string             `json:"mac"`
	Ip           string             `json:"ip"`
	Capabilities []Capability       `json:"capabilities" bson:"capabilities"`
	Zone         Zone               `json:"zone" bson:"zone,omitempty"`
}
