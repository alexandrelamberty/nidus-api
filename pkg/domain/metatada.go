package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Metadata struct {
	DeviceID primitive.ObjectID `json:"device_id" bson:"device_id,omitempty"`
	Type     string             `json:"type" bson:"type"`
}
