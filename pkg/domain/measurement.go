package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Measurement struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	device    Device             `json:"device" bson:"device,omitempty"`
	startDate primitive.DateTime `json:"start_date" bson:"start_date"`
	endDate   primitive.DateTime `json:"date_date" bson:"date_date"`
}
