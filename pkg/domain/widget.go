package domain

type Widget struct {
	Name     string `json:"name" bson:"name,omitempty"`
	DeviceID string `json:"device_id" bson:"device_id,omitempty"`
	// FIXME: capability ?
	Type    string `json:"type" bson:"type,omitempty"`
	Enabled bool   `json:"enabled" bson:"enabled,omitempty"`
}
