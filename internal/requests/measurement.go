package requests

import (
	"nidus-server/pkg/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateMeasurementRequest struct {
	Metadata  domain.Metadata    `json:"metadata" bson:"metadata,omitempty"`
	Value     float32            `json:"value" bson:"value"`
	Timestamp primitive.DateTime `json:"timestamp" bson:"timestamp"`
}
