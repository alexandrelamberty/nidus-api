package requests

type CreateCapability struct {
	Type string `json:"type" validate:"required,min=4,max=32"`
	Kind string `json:"kind" validate:"required,min=4,max=32"`
}
