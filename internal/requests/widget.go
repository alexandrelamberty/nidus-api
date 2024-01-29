package requests

type CreateWidget struct {
	Name     string `json:"name" validate:"required,min=4,max=32"`
	DeviceID string `json:"device_id" validate:"required,hexadecimal,len=24"`
	// FIXME: capability ?
	Type    string `json:"type" validate:"required,min=4,max=32"`
	Enabled bool   `json:"enabled" validate:"required, boolean"`
}
