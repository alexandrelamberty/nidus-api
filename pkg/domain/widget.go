package domain

type Widget struct {
	Name     string `json:"name" bson:"name,omitempty"`
	DeviceID string `json:"device_id" bson:"device_id,omitempty"`
	Type     string `json:"type" bson:"type,omitempty"`
}
