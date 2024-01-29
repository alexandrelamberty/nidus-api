package requests

import (
	"nidus-server/pkg/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateMeasurementRequest struct {
	Metadata  domain.Metadata    `json:"metadata" validate:"required"`
	Value     float32            `json:"value" validate:"required, number"`
	Timestamp primitive.DateTime `json:"timestamp" validate:"required"`
}
