package requests

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Metadata struct {
	DeviceID primitive.ObjectID `json:"device_id,omitempty" validate:"required,mongodb"`
	Type     string             `json:"type" validate:"required,min=4,max=32"`
}

type CreateMeasurementRequest struct {
	Metadata  Metadata           `json:"metadata" validate:"required"`
	Value     float32            `json:"value" validate:"required, number"`
	Timestamp primitive.DateTime `json:"timestamp" validate:"required"`
}
