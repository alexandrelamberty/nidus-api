package requests

type CreateZone struct {
	Name  string `json:"name" validate:"required,min=4,max=32"`
	Floor string `json:"floor" validate:"required,min=3,max=32"`
}
