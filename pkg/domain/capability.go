package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Capability represents a capability of an IoT device.
// It is made up of the type of capacity, ie: sensor, controller and the type of capacity,
// ie: temperature, switch
type Capability struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type string             `json:"type" bson:"type,omitempty"`
	Kind string             `json:"kind" bson:"kind,omitempty"`
}
