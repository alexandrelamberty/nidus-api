package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Role string             `json:"role"`
}
