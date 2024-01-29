package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID             primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name           string               `json:"name" bson:"name,omitempty"`
	Mac            string               `json:"mac" bson:"mac,omitempty"`
	Ip             string               `json:"ip" bson:"ip,omitempty"`
	CapabilityIDs  []primitive.ObjectID `json:"capability_ids" bson:"capability_ids,omitempty"`
	Capabilities   []*Capability        `json:"capabilities" bson:"-",omitempty`
	ZoneID         primitive.ObjectID   `json:"zone_id" bson:"zone_id,omitempty"`
	Zone           *Zone                `json:"zone" bson:"-",omitempty`
	Paired         bool                 `json:"paired" bson:"paired,omitempty"`
	Zoned          bool                 `json:"zoned" bson:"zoned,omitempty"`
	LastConnection string               `json:"last_connection" bson:"last_connection,omitempty"`
}
