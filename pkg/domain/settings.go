package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Settings struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name,omitempty"`
	APIKey string             `json:"api_key" bson:"api_key,omitempty"`
}
