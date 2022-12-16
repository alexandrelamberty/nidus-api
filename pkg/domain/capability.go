package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Capability struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Type string             `json:"type"`
	Kind string             `json:"kind"`
}
