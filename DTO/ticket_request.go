package DTO

type RaiseRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Title  string `json:"title" validate:"required,min=20"`
	Issue  string `json:"issue" validate:"required,min=50"`
}

type GetRequest struct {
	UserID string `json:"user_id" validate:"required"`
}
