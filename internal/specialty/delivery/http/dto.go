package http

import "github.com/google/uuid"

type CreateSpecialtyRequest struct {
	Code string `json:"code" validate:"required,min=2,max=50"`
	Name string `json:"name" validate:"required,min=3,max=255"`
}

type UpdateSpecialtyRequest struct {
	Code string `json:"code" validate:"required,min=2,max=50"`
	Name string `json:"name" validate:"required,min=3,max=255"`
}

type SpecialtyResponse struct {
	ID        uuid.UUID `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt string    `json:"created_at"`
}
