package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Measurement struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Metadata  Metadata           `json:"metadata" bson:"metadata,omitempty"`
	Value     float32            `json:"value" bson:"value"`
	Timestamp primitive.DateTime `json:"timestamp" bson:"timestamp"`
}
