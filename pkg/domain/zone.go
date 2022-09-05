package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Zone struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty"`
}
