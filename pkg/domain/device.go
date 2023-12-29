package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name           string             `json:"name"`
	Mac            string             `json:"mac"`
	Ip             string             `json:"ip"`
	Capabilities   []Capability       `json:"capabilities" bson:"capabilities,omitempty"`
	Zone           Zone               `json:"zone" bson:"zone,omitempty"`
	Paired         string             `json:"paired" bson:"paired,omitempty"`
	LastConnection string             `json:"last_connection" bson:"last_connection,omitempty"`
	Measurements   string             `json:"measurements" bson:"measurements,omitempty"`
}
