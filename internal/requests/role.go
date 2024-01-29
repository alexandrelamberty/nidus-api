package requests

type CreateRole struct {
	Role string `json:"role" validate:"required"`
}

type UpdateRole struct {
	Role string `json:"role" validate:"required"`
}
